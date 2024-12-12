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

	score := ScoreGrid(grid, ScoreFence)
	assert.Equal(t, score, 140)
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreFence)
	assert.Equal(t, score, 772)
}

func TestSimple3Part1(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreFence)
	assert.Equal(t, score, 1930)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreFence)
	assert.Equal(t, score, 1370258)
}

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 80)
}

func TestSimple2Part2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 436)
}

func TestSimple3Part2(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 1206)
}

func TestSimple4Part2(t *testing.T) {
	content := lib.ReadInput("test4.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 236)
}

func TestSimple5Part2(t *testing.T) {
	content := lib.ReadInput("test5.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 368)
}

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	grid := lib.ParseObjectGridFromRunes(content, CreatePlot)
	MarkGrid(&grid)

	score := ScoreGrid(grid, ScoreCorner)
	assert.Equal(t, score, 805814)
}
