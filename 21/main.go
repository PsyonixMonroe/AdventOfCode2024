package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

func GetBestNumeric(code []rune) []rune {
	firstPads := GetAllNumeric(lib.Location{X: 3, Y: 2}, code)
	bestPath := []rune{}
	bestPathLen := int64(0)
	for _, firstPad := range firstPads {
		padLen := RunIterations(firstPad, 25)
		if len(bestPath) == 0 || padLen < bestPathLen {
			bestPath = firstPad
			bestPathLen = padLen
		}
	}

	return bestPath
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

func GetProduction() map[string][]string {
	// <v better than v<
	// <^ better than ^<
	// v> better than >v
	// ^> better than >^

	cache := make(map[string][]string)
	cache["A"] = []string{"A"}
	cache["^A"] = []string{"<A", ">A"}
	cache["^^A"] = []string{"<A", "A", ">A"}
	cache["^^^A"] = []string{"<A", "A", "A", ">A"}
	cache["^<<A"] = []string{"<A", "v<A", "A", ">>^A"}
	cache["^^<<A"] = []string{"<A", "A", "v<A", "A", ">>^A"}

	cache["^>A"] = []string{"<A", "v>A", "^A"}
	cache["^^>A"] = []string{"<A", "A", "v>A", "^A"}
	cache["^^>>A"] = []string{"<A", "A", "v>A", "A", "^A"}

	cache[">A"] = []string{"vA", "^A"}
	cache[">>A"] = []string{"vA", "A", "^A"}

	cache["<A"] = []string{"v<<A", ">>^A"}
	cache["<^A"] = []string{"v<<A", ">^A", ">A"}
	cache["<^^A"] = []string{"v<<A", ">^A", "A", ">A"}

	cache["<vA"] = []string{"v<<A", ">A", "^>A"}

	cache["<<A"] = []string{"v<<A", "A", ">>^A"}
	cache["<<^A"] = []string{"v<<A", "A", ">^A", ">A"}
	cache["<<^^A"] = []string{"v<<A", "A", ">^A", "A", ">A"}

	cache[">^A"] = []string{"vA", "<^A", ">A"}
	cache[">^^A"] = []string{"vA", "<^A", "A", ">A"}
	cache[">>^A"] = []string{"vA", "A", "<^A", ">A"}
	cache[">>^^A"] = []string{"vA", "A", "<^A", "A", ">A"}

	cache[">vvA"] = []string{"vA", "<A", "A", "^>A"}
	cache[">>vvA"] = []string{"vA", "A", "<A", "A", "^>A"}

	cache["vA"] = []string{"<vA", "^>A"}
	cache["vvA"] = []string{"<vA", "A", "^>A"}
	cache["vvvA"] = []string{"<vA", "A", "A", "^>A"}

	cache["v<A"] = []string{"<vA", "<A", ">>^A"}
	cache["v<<A"] = []string{"<vA", "<A", "A", ">>^A"}

	cache["v>A"] = []string{"<vA", ">A", "^A"}
	cache["vv>A"] = []string{"<vA", "A", ">A", "^A"}
	cache["vv>>A"] = []string{"<vA", "A", ">A", "A", "^A"}

	return cache
}

func RunIterations(path []rune, iterations int) int64 {
	production := GetProduction()
	current := make(map[string]int64)
	splits := strings.Split(string(path), "A")
	for _, s := range splits[:len(splits)-1] {
		move := s + "A"
		curCount, found := current[move]
		if !found {
			curCount = 0
		}
		current[move] = curCount + 1
	}

	for range iterations {
		newLevel := make(map[string]int64)

		for k, v := range current {
			produced, found := production[k]
			if !found {
				fmt.Fprintf(os.Stderr, "Unknown pattern found: %s\n", k)
				return -1
			}
			for _, s := range produced {
				newLevelCount, found := newLevel[s]
				if !found {
					newLevelCount = 0
				}
				newLevel[s] = newLevelCount + v
			}
		}

		current = newLevel
	}

	sum := int64(0)
	for k, v := range current {
		sum += v * int64(len(k))
	}

	return sum
}

func GetInputForCode25(code []rune) int64 {
	bestPath := GetBestNumeric(code)
	return RunIterations(bestPath, 25)
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
	updo := []rune{}
	leri := []rune{}
	for range lib.Abs(diff.X) {
		if diff.X > 0 {
			updo = append(updo, 'v')
		} else {
			updo = append(updo, '^')
		}
	}

	for range lib.Abs(diff.Y) {
		if diff.Y > 0 {
			leri = append(leri, '>')
		} else {
			leri = append(leri, '<')
		}
	}

	move := []rune{}
	if diff.X == 0 {
		move = leri
	}
	if diff.Y == 0 {
		move = updo
	}

	nogo := lib.Location{X: 3, Y: 0}
	if diff.X > 0 && diff.Y > 0 {
		// down-right
		move = append(updo, leri...)
		if (nogo.Equal(lib.Location{X: current.X + diff.X, Y: current.Y})) {
			move = append(leri, updo...)
		}
	}
	if diff.X > 0 && diff.Y < 0 {
		// down-left
		move = append(leri, updo...)
		if (nogo.Equal(lib.Location{X: current.X, Y: current.Y + diff.Y})) {
			move = append(updo, leri...)
		}
	}
	if diff.X < 0 && diff.Y > 0 {
		// up-right
		move = append(updo, leri...)
	}
	if diff.X < 0 && diff.Y < 0 {
		// up-left
		move = append(leri, updo...)
		// todo
		if (nogo.Equal(lib.Location{X: current.X, Y: current.Y + diff.Y})) {
			move = append(updo, leri...)
		}
	}

	move = append(move, 'A')

	if len(subpaths) == 0 {
		paths = append(paths, move)
	}
	for _, subpath := range subpaths {
		paths = append(paths, append(move, subpath...))
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

func GetComplexityForCode(code []rune, thirdPad int64) int64 {
	strCode := string(code[:len(code)-1])
	parsedCode, err := strconv.Atoi(strCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse %s\n", strCode)
		return -1
	}

	return int64(parsedCode) * thirdPad
}
