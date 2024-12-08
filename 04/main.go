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
	// fmt.Fprintf(os.Stderr, "[%d, %d] - Columns: %d\n", row, column, c.columns)
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
				// fmt.Fprintf(os.Stderr, "[%d, %d - %s]: UpdatedCount: %d\n", i, j, string(current), xmasCount)
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

	// fmt.Fprintf(os.Stderr, "[%d, %d]: Found Counts: %d\n", row, column, count)

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

	// fmt.Fprintf(os.Stderr, "Up %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "UpRight %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "Right %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "DownRight %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}

func CheckVerticalDown(crossword CrossWord, row int, column int) bool {
	if row >= (crossword.rows - 3) {
		return false
	}
	// fmt.Fprintf(os.Stderr, "Down [%d, %d] - Max: %d\n", row, column, crossword.rows)
	expectX := crossword.Get(row, column)
	expectM := crossword.Get(row+1, column)
	expectA := crossword.Get(row+2, column)
	expectS := crossword.Get(row+3, column)

	// fmt.Fprintf(os.Stderr, "Down %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "DownLeft %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "Left %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

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

	// fmt.Fprintf(os.Stderr, "UpLeft %s %s %s %s\n", string(crossword.Get(row, column)), string(expectM), string(expectA), string(expectS))

	return expectX == 'X' && expectM == 'M' && expectA == 'A' && expectS == 'S'
}
