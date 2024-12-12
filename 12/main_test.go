package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid)
	assert.Equal(t, score, 140)
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid)
	assert.Equal(t, score, 772)
}

func TestSimple3Part1(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid)
	assert.Equal(t, score, 1930)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid)
	assert.Equal(t, score, 1370258)
}
