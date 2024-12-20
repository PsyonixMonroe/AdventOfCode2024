package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	track := ParseInput(content)
	sum := CountGoodPaths(track, content, 64)
	assert.Equal(t, sum, 1)

	sum = CountGoodPaths(track, content, 19)
	assert.Equal(t, sum, 5)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	track := ParseInput(content)
	sum := CountGoodPaths(track, content, 100)
	assert.Equal(t, sum, 1375)
}
