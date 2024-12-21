package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type KeyPad struct {
	cursor lib.Location
	grid   lib.Grid[rune]
}

func (k KeyPad) GetKeyLocation(r rune) lib.Location {
	for i := range k.grid.Rows() {
		for j := range k.grid.Columns() {
			if r == *k.grid.Get(i, j).Get() {
				return lib.Location{X: i, Y: j}
			}
		}
	}

	return lib.Location{X: -1, Y: -1}
}

func GetNumericGrid() lib.Grid[rune] {
	g := lib.GetGrid[rune](4, 3)
	g.Put('7', 0, 0)
	g.Put('8', 0, 1)
	g.Put('9', 0, 2)
	g.Put('4', 1, 0)
	g.Put('5', 1, 1)
	g.Put('6', 1, 2)
	g.Put('1', 2, 0)
	g.Put('2', 2, 1)
	g.Put('3', 2, 2)
	g.Put('0', 3, 1)
	g.Put('A', 3, 2)

	g.Put('X', 3, 0)

	return g
}

func GetArrowGrid() lib.Grid[rune] {
	g := lib.GetGrid[rune](4, 3)
	g.Put('^', 0, 1)
	g.Put('A', 0, 2)
	g.Put('<', 1, 0)
	g.Put('v', 1, 1)
	g.Put('>', 1, 2)

	g.Put('X', 0, 0)

	return g
}

func getCounts(raw []rune) (int, int, int, int) {
	numLeft := 0
	numRight := 0
	numDown := 0
	numUp := 0
	for _, key := range raw {
		switch key {
		case '<':
			numLeft++
		case '>':
			numRight++
		case '^':
			numUp++
		case 'v':
			numDown++
		}
	}

	return numLeft, numRight, numDown, numUp
}

func SortDirection(raw []rune) []rune {
	numLeft, numRight, numDown, numUp := getCounts(raw)
	output := []rune{}
	for range numDown {
		output = append(output, 'v')
	}
	for range numLeft {
		output = append(output, '<')
	}
	for range numRight {
		output = append(output, '>')
	}
	for range numUp {
		output = append(output, '^')
	}

	return output
}

func SortNumericOptimal(raw []rune, cursor lib.Location) []rune {
	if cursor.X < 3 {
		numLeft, numRight, numDown, numUp := getCounts(raw)
		output := []rune{}
		for range numLeft {
			output = append(output, '<')
		}
		for range numRight {
			output = append(output, '>')
		}
		for range numUp {
			output = append(output, '^')
		}
		for range numDown {
			output = append(output, 'v')
		}
		return output
	} else {
		return SortNumericSimple(raw)
	}
}

func SortNumericSimple(raw []rune) []rune {
	numLeft, numRight, numDown, numUp := getCounts(raw)
	output := []rune{}
	for range numUp {
		output = append(output, '^')
	}
	for range numLeft {
		output = append(output, '<')
	}

	for range numRight {
		output = append(output, '>')
	}
	for range numDown {
		output = append(output, 'v')
	}

	return output
}

func TranslateFromNumeric(keypress []rune) []rune {
	pad := GetNumericGrid()
	keypad := KeyPad{cursor: lib.Location{X: 3, Y: 2}, grid: pad}
	output := []rune{}
	for _, key := range keypress {
		moves := GetMovesToKey(keypad, key)
		moves = append(SortNumericOptimal(moves, keypad.cursor), 'A')
		keypad.cursor = keypad.GetKeyLocation(key)
		output = append(output, moves...)
	}

	return output
}

func TranslateFromArrowPad(keypress []rune) []rune {
	pad := GetArrowGrid()
	keypad := KeyPad{cursor: lib.Location{X: 0, Y: 2}, grid: pad}

	output := []rune{}
	for _, key := range keypress {
		moves := GetMovesToKey(keypad, key)
		keypad.cursor = keypad.GetKeyLocation(key)
		moves = append(SortDirection(moves), 'A')
		output = append(output, moves...)
	}
	return output
}

func GetMovesToKey(keypad KeyPad, key rune) []rune {
	current := keypad.cursor
	final := keypad.GetKeyLocation(key)
	moves := []rune{}
	for !current.Equal(final) {
		direction := GetUnitDirection(lib.Location{X: final.X - current.X, Y: final.Y - current.Y})
		if direction.X == 0 || keypad.GetKeyLocation('X').Equal(lib.Location{X: current.X + direction.X, Y: current.Y}) {
			// try to move in the Y direction
			if direction.Y > 0 {
				current.Y = current.Y + 1
				moves = append(moves, '>')
			} else {
				current.Y = current.Y - 1
				moves = append(moves, '<')
			}
		} else {
			// try to move in the X direction
			if direction.X > 0 {
				current.X = current.X + 1
				moves = append(moves, 'v')
			} else {
				current.X = current.X - 1
				moves = append(moves, '^')
			}
		}
	}

	return moves
}

func GetUnitDirection(scaleDir lib.Location) lib.Location {
	newLoc := lib.Location{X: 0, Y: 0}

	if scaleDir.X > 0 {
		newLoc.X = 1
	} else if scaleDir.X < 0 {
		newLoc.X = -1
	}

	if scaleDir.Y > 0 {
		newLoc.Y = 1
	} else if scaleDir.Y < 0 {
		newLoc.Y = -1
	}

	return newLoc
}

func GetInputForCode(code []rune) []rune {
	firstPads := GetAllNumeric(lib.Location{X: 3, Y: 2}, code)
	bestPath := []rune{}
	for _, firstPad := range firstPads {
		secondPad := TranslateFromArrowPad(firstPad)
		thirdPad := TranslateFromArrowPad(secondPad)
		if len(bestPath) == 0 || len(thirdPad) < len(bestPath) {
			bestPath = thirdPad
		}
	}

	return bestPath
}

func GetVertPath(diff lib.Location) []rune {
	output := []rune{}
	for diff.X != 0 {
		if diff.X < 0 {
			output = append(output, '^')
			diff.X++
		} else {
			output = append(output, 'v')
			diff.X--
		}
	}

	for diff.Y != 0 {
		if diff.Y < 0 {
			output = append(output, '<')
			diff.Y++
		} else {
			output = append(output, '>')
			diff.Y--
		}
	}

	output = append(output, 'A')

	return output
}

func GetHorzPath(diff lib.Location) []rune {
	output := []rune{}
	for diff.Y != 0 {
		if diff.Y < 0 {
			output = append(output, '<')
			diff.Y++
		} else {
			output = append(output, '>')
			diff.Y--
		}
	}

	for diff.X != 0 {
		if diff.X < 0 {
			output = append(output, '^')
			diff.X++
		} else {
			output = append(output, 'v')
			diff.X--
		}
	}

	output = append(output, 'A')

	return output
}

func HorizViable(diff lib.Location, current lib.Location) bool {
	return current.X != 3 || current.Y == 2 && diff.Y == -1 || current.Y == 1 && diff.Y == 0
}

func VertViable(diff lib.Location, current lib.Location) bool {
	return current.Y != 0 || current.X == 2 && diff.X == 0 || current.X == 1 && diff.X == 1 || current.X == 0 && (diff.X == 1 || diff.X == 2)
}

func GetAllNumeric(current lib.Location, code []rune) [][]rune {
	paths := [][]rune{}
	if len(code) == 0 {
		return paths
	}
	grid := GetNumericGrid()
	button := code[0]
	target := GetKeyLocation(grid, button)
	subpaths := GetAllNumeric(target, code[1:])
	diff := lib.Location{X: target.X - current.X, Y: target.Y - current.Y}
	horizPath := []rune{}
	vertPath := []rune{}
	if HorizViable(diff, current) {
		horizPath = GetHorzPath(diff)
	}
	if VertViable(diff, current) {
		vertPath = GetVertPath(diff)
	}

	if (len(horizPath) != 0 && string(horizPath) != string(vertPath)) || len(vertPath) == 0 {
		if len(subpaths) == 0 {
			paths = append(paths, horizPath)
		}
		for _, subpath := range subpaths {
			paths = append(paths, append(horizPath, subpath...))
		}
	}

	if len(vertPath) != 0 {
		if len(subpaths) == 0 {
			paths = append(paths, vertPath)
		}
		for _, subpath := range subpaths {
			paths = append(paths, append(vertPath, subpath...))
		}
	}

	return paths
}

func GetKeyLocation(k lib.Grid[rune], r rune) lib.Location {
	for i := range k.Rows() {
		for j := range k.Columns() {
			if r == *k.Get(i, j).Get() {
				return lib.Location{X: i, Y: j}
			}
		}
	}

	return lib.Location{X: -1, Y: -1}
}

func GetComplexityForCode(code []rune, thirdPad []rune) int {
	strCode := string(code[:len(code)-1])
	parsedCode, err := strconv.Atoi(strCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse %s\n", strCode)
		return -1
	}

	return parsedCode * len(thirdPad)
}
