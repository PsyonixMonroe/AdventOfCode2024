package main

import (
	"fmt"
	"os"
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

type Location struct {
	station rune
	points  []rune
}

type Map struct {
	locs    []Location
	rows    int
	columns int
}

func (m Map) loc(row int, column int) int {
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

func ParseInput(content string) Map {
	locs := []Location{}
	rows := strings.Split(content, "\n")
	numRows := len(rows)
	if numRows == 0 {
		fmt.Fprintf(os.Stderr, "No Rows Found\n")
		return Map{}
	}
	numColumns := len(rows[0])

	for _, val := range rows {
		row := strings.Trim(val, "\n")
		if row == "" {
			numRows--
			continue
		}
		for _, char := range row {
			locs = append(locs, Location{station: char, points: []rune{}})
		}
	}

	return Map{locs: locs, rows: numRows, columns: numColumns}
}

func MarkPoints(m *Map) {
	for i := range m.rows {
		for j := range m.columns {
			loc := m.locs[m.loc(i, j)]
			if loc.station == '.' {
				continue
			}
			MarkStationPoints(m, loc.station, i, j)
		}
	}
}

func MarkStationPoints(m *Map, originStation rune, originX int, originY int) {
	for i := range m.rows {
		for j := range m.columns {
			loc := m.locs[m.loc(i, j)]
			if loc.station == originStation {
				if i == originX && j == originY {
					continue
				}
				// compute and mark points
				var x1, y1, x2, y2 int
				if originX < i {
					x1, y1, x2, y2 = IntersectionsFromStation(originX, originY, i, j)
				} else {
					x1, y1, x2, y2 = IntersectionsFromStation(i, j, originX, originY)
				}

				if x1 >= 0 && x1 < m.rows && y1 >= 0 && y1 < m.columns {
					m.locs[m.loc(x1, y1)].points = append(m.locs[m.loc(x1, y1)].points, originStation)
				}
				if x2 >= 0 && x2 < m.rows && y2 >= 0 && y2 < m.columns {
					m.locs[m.loc(x2, y2)].points = append(m.locs[m.loc(x2, y2)].points, originStation)
				}
			}
		}
	}
}

func GetRiseAndRun(upperX int, upperY int, lowerX int, lowerY int) (int, int) {
	rise := lowerX - upperX
	run := lowerY - upperY
	return rise, run
}

func IntersectionsFromStation(upperX int, upperY int, lowerX int, lowerY int) (int, int, int, int) {
	rise, run := GetRiseAndRun(upperX, upperY, lowerX, lowerY)
	highX := upperX - rise
	lowX := lowerX + rise
	highY := upperY - run
	lowY := lowerY + run
	return highX, highY, lowX, lowY
}

func CountLocations(m Map) int {
	sum := 0
	for i := range m.rows {
		for j := range m.columns {
			loc := m.locs[m.loc(i, j)]
			if len(loc.points) > 0 {
				sum++
			}
		}
	}

	return sum
}
