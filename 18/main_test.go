package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	// test parameters
	gridRows := 7
	gridColumns := 7
	numBytes := 12

	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	grid, byteLocations := ParseBytes(content, gridRows, gridColumns)
	assert.NotEqual(t, len(byteLocations), 0)

	start := lib.Location{X: 0, Y: 0}
	end := lib.Location{X: gridRows - 1, Y: gridColumns - 1}
	path := FindPath(grid, start, end, byteLocations, numBytes)
	assert.Equal(t, len(path), 22)
}

func TestFullPart1(t *testing.T) {
	// test parameters
	gridRows := 71
	gridColumns := 71
	numBytes := 1024

	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	grid, byteLocations := ParseBytes(content, gridRows, gridColumns)
	assert.NotEqual(t, len(byteLocations), 0)

	start := lib.Location{X: 0, Y: 0}
	end := lib.Location{X: gridRows - 1, Y: gridColumns - 1}
	path := FindPath(grid, start, end, byteLocations, numBytes)
	assert.Equal(t, len(path), 322)
}
