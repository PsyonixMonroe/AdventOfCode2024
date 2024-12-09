package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	DirUp    = iota
	DirDown  = iota
	DirLeft  = iota
	DirRight = iota
)

type MapSimple struct {
	room      []rune
	rows      int
	columns   int
	cursorX   int
	cursorY   int
	cursorDir int
}

func (m MapSimple) loc(row int, column int) int {
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

func (m MapSimple) HasExitedRoom() bool {
	return m.cursorX < 0 ||
		m.cursorY < 0 ||
		m.cursorX >= m.columns ||
		m.cursorY >= m.rows
}

func (m MapSimple) IsBlocked(row int, column int) bool {
	if row < 0 || column < 0 || column >= m.columns || row >= m.rows {
		// we need to allow movement OOB to determine HasExitedRoom()
		// but don't want to call loc() for those locations
		return false
	}
	return m.room[m.loc(row, column)] == '#'
}

func (m *MapSimple) RotateCursor() {
	switch m.cursorDir {
	case DirUp:
		m.cursorDir = DirRight
		return
	case DirDown:
		m.cursorDir = DirLeft
		return
	case DirLeft:
		m.cursorDir = DirUp
		return
	case DirRight:
		m.cursorDir = DirDown
		return
	}
}

func (m MapSimple) MarkLocation() {
	m.room[m.loc(m.cursorX, m.cursorY)] = 'X'
}

func (m *MapSimple) MoveCursor() {
	switch m.cursorDir {
	case DirUp:
		if m.IsBlocked(m.cursorX-1, m.cursorY) {
			m.RotateCursor()
		} else {
			m.MarkLocation()
			m.cursorX -= 1
		}
	case DirDown:
		if m.IsBlocked(m.cursorX+1, m.cursorY) {
			m.RotateCursor()
		} else {
			m.MarkLocation()
			m.cursorX += 1
		}
	case DirLeft:
		if m.IsBlocked(m.cursorX, m.cursorY-1) {
			m.RotateCursor()
		} else {
			m.MarkLocation()
			m.cursorY -= 1
		}
	case DirRight:
		if m.IsBlocked(m.cursorX, m.cursorY+1) {
			m.RotateCursor()
		} else {
			m.MarkLocation()
			m.cursorY += 1
		}
	}
}

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func ParseInput(content string) MapSimple {
	room := []rune{}
	rows := strings.Split(content, "\n")
	numRows := len(rows)
	if numRows == 0 {
		fmt.Fprintf(os.Stderr, "No Rows Found\n")
		return MapSimple{}
	}
	numColumns := len(rows[0])
	cursorX := 0
	cursorY := 0
	cursorDir := DirUp

	for i, val := range rows {
		row := strings.Trim(val, "\n")
		if row == "" {
			numRows--
			continue
		}
		for j, char := range row {
			if char == '^' {
				cursorX = i
				cursorY = j
				cursorDir = DirUp
			}
			room = append(room, char)
		}
	}

	return MapSimple{room: room, rows: numRows, columns: numColumns, cursorX: cursorX, cursorY: cursorY, cursorDir: cursorDir}
}

func WalkMap(room *MapSimple) {
	for {
		if room.HasExitedRoom() {
			break
		}
		room.MoveCursor()
	}
}

func CountPath(room MapSimple) int {
	count := 0
	for _, val := range room.room {
		if val == 'X' {
			count++
		}
	}

	return count
}

type Location struct {
	obj     rune
	moveDir []int
}

type MapComplex struct {
	room      []Location
	rows      int
	columns   int
	cursorX   int
	cursorY   int
	cursorDir int
}

func (m MapComplex) Print() {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Rows: %d\nColumns:%d\nCursor: [%d, %d] - %d\nRoom Size:%d\n", m.rows, m.columns, m.cursorX, m.cursorY, m.cursorDir, len(m.room)))
	for i, loc := range m.room {
		if i%m.columns == 0 {
			sb.WriteString("\n")
		}
		sb.WriteRune(loc.obj)
	}
	sb.WriteString("\n")
	fmt.Fprint(os.Stderr, sb.String())
}

func (m MapComplex) loc(row int, column int) int {
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

func (m MapComplex) HasExitedRoom() bool {
	return m.cursorX < 0 ||
		m.cursorY < 0 ||
		m.cursorX >= m.columns ||
		m.cursorY >= m.rows
}

func (m MapComplex) IsBlocked(row int, column int) bool {
	if row < 0 || column < 0 || column >= m.columns || row >= m.rows {
		// we need to allow movement OOB to determine HasExitedRoom()
		// but don't want to call loc() for those locations
		return false
	}
	return m.room[m.loc(row, column)].obj == '#'
}

func (m *MapComplex) RotateCursor() {
	switch m.cursorDir {
	case DirUp:
		m.cursorDir = DirRight
		return
	case DirDown:
		m.cursorDir = DirLeft
		return
	case DirLeft:
		m.cursorDir = DirUp
		return
	case DirRight:
		m.cursorDir = DirDown
		return
	}
}

func (m MapComplex) MarkLocation(direction int) {
	loc := &(m.room[m.loc(m.cursorX, m.cursorY)])
	loc.obj = 'X'
	loc.moveDir = append(loc.moveDir, direction)
}

func (m MapComplex) HasLooped() bool {
	location := m.room[m.loc(m.cursorX, m.cursorY)]
	for _, prevDir := range location.moveDir {
		if m.cursorDir == prevDir {
			// we have been here before
			return true
		}
	}

	return false
}

func (m *MapComplex) MoveCursor() {
	switch m.cursorDir {
	case DirUp:
		if m.IsBlocked(m.cursorX-1, m.cursorY) {
			m.RotateCursor()
		} else {
			m.MarkLocation(DirUp)
			m.cursorX -= 1
		}
	case DirDown:
		if m.IsBlocked(m.cursorX+1, m.cursorY) {
			m.RotateCursor()
		} else {
			m.MarkLocation(DirDown)
			m.cursorX += 1
		}
	case DirLeft:
		if m.IsBlocked(m.cursorX, m.cursorY-1) {
			m.RotateCursor()
		} else {
			m.MarkLocation(DirLeft)
			m.cursorY -= 1
		}
	case DirRight:
		if m.IsBlocked(m.cursorX, m.cursorY+1) {
			m.RotateCursor()
		} else {
			m.MarkLocation(DirRight)
			m.cursorY += 1
		}
	}
}

func (m MapComplex) copy() MapComplex {
	roomCopy := make([]Location, len(m.room))
	copy(roomCopy, m.room)
	return MapComplex{room: roomCopy, rows: m.rows, columns: m.columns, cursorX: m.cursorX, cursorY: m.cursorY, cursorDir: m.cursorDir}
}

func WalkAllMaps(m MapComplex) int {
	count := 0
	for i := range m.rows {
		for j := range m.columns {
			char := m.room[m.loc(i, j)].obj
			if char == '#' || char == '^' {
				// location is start or already blocked, continue
				continue
			}
			testCopy := m.copy()
			testCopy.room[testCopy.loc(i, j)].obj = '#'
			hasLoop := SimulateLoop(testCopy)
			if hasLoop {
				count++
			}
		}
	}

	return count
}

func SimulateLoop(m MapComplex) bool {
	for {
		if m.HasExitedRoom() {
			return false
		}
		if m.HasLooped() {
			break
		}
		m.MoveCursor()
	}

	return true
}

func ParseInputComplex(content string) MapComplex {
	room := []Location{}
	rows := strings.Split(content, "\n")
	numRows := len(rows)
	if numRows == 0 {
		fmt.Fprintf(os.Stderr, "No Rows Found\n")
		return MapComplex{}
	}
	numColumns := len(rows[0])
	cursorX := 0
	cursorY := 0
	cursorDir := DirUp

	for i, val := range rows {
		row := strings.Trim(val, "\n")
		if row == "" {
			numRows--
			continue
		}
		for j, char := range row {
			if char == '^' {
				cursorX = i
				cursorY = j
				cursorDir = DirUp
			}
			room = append(room, Location{obj: char, moveDir: []int{}})
		}
	}

	return MapComplex{room: room, rows: numRows, columns: numColumns, cursorX: cursorX, cursorY: cursorY, cursorDir: cursorDir}
}
