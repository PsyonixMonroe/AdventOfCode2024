package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

type TopoMap struct {
	grid    []int
	rows    int
	columns int
}

func (m TopoMap) loc(row int, column int) int {
	if row < 0 {
		fmt.Fprintf(os.Stderr, "Unable to find loc, row is negative: %d\n", row)
		return -1
	}
	if column < 0 {
		fmt.Fprintf(os.Stderr, "Unable to find loc, column is negative: %d\n", column)
		return -1
	}
	if column >= m.columns {
		fmt.Fprintf(os.Stderr, "Unable to find loc, column is OOB: %d\n", column)
		return -1
	}
	if row >= m.rows {
		fmt.Fprintf(os.Stderr, "Unable to find loc, row is OOB: %d\n", row)
		return -1
	}
	return row*m.columns + column
}

func ParseTopoMap(content string) TopoMap {
	lines := strings.Split(strings.Trim(content, "\n"), "\n")
	rowCount := len(lines)
	if rowCount == 0 {
		return TopoMap{grid: []int{}, rows: 0, columns: 0}
	}
	columnCount := len(strings.Trim(lines[0], "\n"))
	grid := []int{}
	for _, line := range lines {
		if line == "" {
			rowCount--
			continue
		}
		// parse row
		for _, char := range line {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse height from char: %s\n", string(char))
				continue
			}
			grid = append(grid, height)
		}
	}

	return TopoMap{grid: grid, rows: rowCount, columns: columnCount}
}

func FindTrails(m TopoMap) int {
	sum := 0

	for i := range m.rows {
		for j := range m.columns {
			if m.grid[m.loc(i, j)] == 0 {
				sum += len(WalkTrail(m, i, j))
			}
		}
	}

	return sum
}

func WalkTrail(m TopoMap, currentRow int, currentColumn int) []int {
	currentHeight := m.grid[m.loc(currentRow, currentColumn)]
	if currentHeight == 9 {
		return []int{m.loc(currentRow, currentColumn)}
	}

	summits := []int{}
	if currentRow > 0 {
		// check up
		nextLoc := m.grid[m.loc(currentRow-1, currentColumn)]
		if nextLoc == currentHeight+1 {
			summits = mergeDistinct(summits, WalkTrail(m, currentRow-1, currentColumn))
		}
	}

	if currentColumn > 0 {
		// check left
		nextLoc := m.grid[m.loc(currentRow, currentColumn-1)]
		if nextLoc == currentHeight+1 {
			summits = mergeDistinct(summits, WalkTrail(m, currentRow, currentColumn-1))
		}
	}

	if currentRow < m.rows-1 {
		// check down
		nextLoc := m.grid[m.loc(currentRow+1, currentColumn)]
		if nextLoc == currentHeight+1 {
			summits = mergeDistinct(summits, WalkTrail(m, currentRow+1, currentColumn))
		}
	}

	if currentColumn < m.columns-1 {
		// check right
		nextLoc := m.grid[m.loc(currentRow, currentColumn+1)]
		if nextLoc == currentHeight+1 {
			summits = mergeDistinct(summits, WalkTrail(m, currentRow, currentColumn+1))
		}
	}

	return summits
}

func FindScoreTrailHeads(m TopoMap) int {
	sum := 0

	for i := range m.rows {
		for j := range m.columns {
			if m.grid[m.loc(i, j)] == 0 {
				sum += len(WalkTrailScore(m, i, j))
			}
		}
	}

	return sum
}

func WalkTrailScore(m TopoMap, currentRow int, currentColumn int) []int {
	currentHeight := m.grid[m.loc(currentRow, currentColumn)]
	if currentHeight == 9 {
		return []int{m.loc(currentRow, currentColumn)}
	}

	summits := []int{}
	if currentRow > 0 {
		// check up
		nextLoc := m.grid[m.loc(currentRow-1, currentColumn)]
		if nextLoc == currentHeight+1 {
			summits = append(summits, WalkTrailScore(m, currentRow-1, currentColumn)...)
		}
	}

	if currentColumn > 0 {
		// check left
		nextLoc := m.grid[m.loc(currentRow, currentColumn-1)]
		if nextLoc == currentHeight+1 {
			summits = append(summits, WalkTrailScore(m, currentRow, currentColumn-1)...)
		}
	}

	if currentRow < m.rows-1 {
		// check down
		nextLoc := m.grid[m.loc(currentRow+1, currentColumn)]
		if nextLoc == currentHeight+1 {
			summits = append(summits, WalkTrailScore(m, currentRow+1, currentColumn)...)
		}
	}

	if currentColumn < m.columns-1 {
		// check right
		nextLoc := m.grid[m.loc(currentRow, currentColumn+1)]
		if nextLoc == currentHeight+1 {
			summits = append(summits, WalkTrailScore(m, currentRow, currentColumn+1)...)
		}
	}

	return summits
}

func mergeDistinct(a []int, b []int) []int {
	merged := make([]int, len(a))
	copy(merged, a)
	for _, val := range b {
		merged = appendDistinct(merged, val)
	}

	return merged
}

func appendDistinct(src []int, newItem int) []int {
	for _, value := range src {
		if value == newItem {
			return src
		}
	}

	return append(src, newItem)
}
