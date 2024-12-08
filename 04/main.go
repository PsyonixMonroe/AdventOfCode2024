package main

import (
	"fmt"
	"os"
	"strings"
)

type CrossWord struct {
	chars   []rune
	columns int
	rows    int
}

func (c CrossWord) Print() {
	sb := strings.Builder{}
	for i, char := range c.chars {
		if i%c.columns == 0 {
			sb.WriteString("\n")
		}
		sb.WriteRune(char)
	}

	fmt.Fprintf(os.Stdout, "Rows: %d\nColumns: %d\n", c.rows, c.columns)
	fmt.Fprintln(os.Stdout, sb.String())
}

func (c CrossWord) loc(row int, column int) int {
	return row*c.columns + column
}

func (c CrossWord) Get(row int, column int) rune {
	return c.chars[c.loc(row, column)]
}

func (c CrossWord) Add(char rune, row int, column int) {
	c.chars[c.loc(row, column)] = char
}

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func ReadCrossword(content string) CrossWord {
	lines := strings.Split(content, "\n")
	characters := make([]rune, len(content))
	crossword := CrossWord{chars: characters, columns: 0, rows: 0}
	for row, line := range lines {
		if line == "" {
			continue
		}
		if crossword.columns == 0 {
			crossword.columns = len(line)
		}
		characters := strings.Trim(line, "\n")
		for column, char := range characters {
			crossword.Add(char, row, column)
		}
		crossword.rows += 1
	}

	return crossword
}

func FindXmas(crossword CrossWord) int {
	xmasCount := 0
	for i := range crossword.rows {
		for j := range crossword.columns {
			current := crossword.Get(i, j)
			if current == 'X' {
				xmasCount += CountAllWays(crossword, i, j)
			}
		}
	}

	return xmasCount
}

func CountAllWays(crossword CrossWord, row int, column int) int {
	count := 0

	if CheckVerticalUp(crossword, row, column) {
		count++
	}
	if CheckDiagonalUpRight(crossword, row, column) {
		count++
	}
	if CheckHorizontalRight(crossword, row, column) {
		count++
	}
	if CheckDiagonalDownRight(crossword, row, column) {
		count++
	}
	if CheckVerticalDown(crossword, row, column) {
		count++
	}
	if CheckDiagonalDownLeft(crossword, row, column) {
		count++
	}
	if CheckHorizontalLeft(crossword, row, column) {
		count++
	}
	if CheckDiagonalUpLeft(crossword, row, column) {
		count++
	}

	return count
}

func CheckVerticalUp(crossword CrossWord, row int, column int) bool {
	if row < 3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row-1, column)
	expectA := crossword.Get(row-2, column)
	expectS := crossword.Get(row-3, column)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckDiagonalUpRight(crossword CrossWord, row int, column int) bool {
	if row < 3 || column > crossword.columns-3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row-1, column+1)
	expectA := crossword.Get(row-2, column+2)
	expectS := crossword.Get(row-3, column+3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckHorizontalRight(crossword CrossWord, row int, column int) bool {
	if column >= crossword.columns-3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row, column+1)
	expectA := crossword.Get(row, column+2)
	expectS := crossword.Get(row, column+3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckDiagonalDownRight(crossword CrossWord, row int, column int) bool {
	if column >= crossword.columns-3 || row >= crossword.rows-3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row+1, column+1)
	expectA := crossword.Get(row+2, column+2)
	expectS := crossword.Get(row+3, column+3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckVerticalDown(crossword CrossWord, row int, column int) bool {
	if row >= (crossword.rows - 3) {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row+1, column)
	expectA := crossword.Get(row+2, column)
	expectS := crossword.Get(row+3, column)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckDiagonalDownLeft(crossword CrossWord, row int, column int) bool {
	if column < 3 || row >= crossword.rows-3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row+1, column-1)
	expectA := crossword.Get(row+2, column-2)
	expectS := crossword.Get(row+3, column-3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckHorizontalLeft(crossword CrossWord, row int, column int) bool {
	if column < 3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row, column-1)
	expectA := crossword.Get(row, column-2)
	expectS := crossword.Get(row, column-3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckDiagonalUpLeft(crossword CrossWord, row int, column int) bool {
	if column < 3 || row < 3 {
		return false
	}
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row-1, column-1)
	expectA := crossword.Get(row-2, column-2)
	expectS := crossword.Get(row-3, column-3)

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func FindXmas2(crossword CrossWord) int {
	count := 0
	for i := range crossword.rows {
		for j := range crossword.columns {
			current := crossword.Get(i, j)
			if current == 'A' {
				found := CheckXMAS(crossword, i, j)
				if found {
					count++
				}
			}
		}
	}

	return count
}

func CheckXMAS(crossword CrossWord, row int, column int) bool {
	checkMASTop := CheckMAS1(crossword, row, column)
	checkMASRight := CheckMAS2(crossword, row, column)
	checkMASBottom := CheckMAS3(crossword, row, column)
	checkMASLeft := CheckMAS4(crossword, row, column)
	return checkMASTop || checkMASRight || checkMASBottom || checkMASLeft
}

func CheckMAS1(crossword CrossWord, row int, column int) bool {
	// M   M
	//	 A
	// S   S
	if row < 1 || row >= crossword.rows-1 {
		return false
	}
	if column < 1 || column >= crossword.columns-1 {
		return false
	}

	topLeft := crossword.Get(row-1, column-1)
	topRight := crossword.Get(row-1, column+1)
	center := crossword.Get(row, column)
	bottomLeft := crossword.Get(row+1, column-1)
	bottomRight := crossword.Get(row+1, column+1)

	return topLeft == 'M' && topRight == 'M' && center == 'A' && bottomLeft == 'S' && bottomRight == 'S'
}

func CheckMAS2(crossword CrossWord, row int, column int) bool {
	// S   S
	//   A
	// M   M
	if row < 1 || row >= crossword.rows-1 {
		return false
	}
	if column < 1 || column >= crossword.columns-1 {
		return false
	}

	topLeft := crossword.Get(row-1, column-1)
	topRight := crossword.Get(row-1, column+1)
	center := crossword.Get(row, column)
	bottomLeft := crossword.Get(row+1, column-1)
	bottomRight := crossword.Get(row+1, column+1)

	return topLeft == 'S' && topRight == 'S' && center == 'A' && bottomLeft == 'M' && bottomRight == 'M'
}

func CheckMAS3(crossword CrossWord, row int, column int) bool {
	// M   S
	//   A
	// M   S
	if row < 1 || row >= crossword.rows-1 {
		return false
	}
	if column < 1 || column >= crossword.columns-1 {
		return false
	}

	topLeft := crossword.Get(row-1, column-1)
	topRight := crossword.Get(row-1, column+1)
	center := crossword.Get(row, column)
	bottomLeft := crossword.Get(row+1, column-1)
	bottomRight := crossword.Get(row+1, column+1)

	return topLeft == 'M' && topRight == 'S' && center == 'A' && bottomLeft == 'M' && bottomRight == 'S'
}

func CheckMAS4(crossword CrossWord, row int, column int) bool {
	// S   M
	//   A
	// S   M
	if row < 1 || row >= crossword.rows-1 {
		return false
	}
	if column < 1 || column >= crossword.columns-1 {
		return false
	}

	topLeft := crossword.Get(row-1, column-1)
	topRight := crossword.Get(row-1, column+1)
	center := crossword.Get(row, column)
	bottomLeft := crossword.Get(row+1, column-1)
	bottomRight := crossword.Get(row+1, column+1)

	return topLeft == 'S' && topRight == 'M' && center == 'A' && bottomLeft == 'S' && bottomRight == 'M'
}
