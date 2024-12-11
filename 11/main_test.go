package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	stones := lib.ParseIntLineFields(content)
	stones = RunStonesSim(stones, 1)
	assertEqual(t, stones, []int{1, 2024, 1, 0, 9, 9, 2021976})
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	stones := lib.ParseIntLineFields(content)
	stones = RunStonesSim(stones, 6)
	assertEqual(t, stones, []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2})
}

func TestSimple2LargePart1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	stones := lib.ParseIntLineFields(content)
	stones = RunStonesSim(stones, 25)
	assert.Equal(t, len(stones), 55312)
}

func assertEqual(t *testing.T, actual []int, expected []int) {
	assert.Equal(t, len(actual), len(expected))
	for i, val := range actual {
		assert.Equal(t, val, expected[i])
	}
}

func TestSplitDigits(t *testing.T) {
	num1, num2 := splitDigits(1234)
	assert.Equal(t, num1, 12)
	assert.Equal(t, num2, 34)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	stones := lib.ParseIntLineFields(content)
	stones = RunStonesSim(stones, 25)
	assert.Equal(t, len(stones), 193269)
}
