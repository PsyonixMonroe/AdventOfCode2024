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

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	track := ParseInput(content)
	var sum int

	sum = CountCheats(track, 76)
	assert.Equal(t, sum, 3)

	sum = CountCheats(track, 74)
	assert.Equal(t, sum, 7)

	sum = CountCheats(track, 72)
	assert.Equal(t, sum, 29)

	sum = CountCheats(track, 70)
	assert.Equal(t, sum, 41)

	sum = CountCheats(track, 68)
	assert.Equal(t, sum, 55)

	sum = CountCheats(track, 66)
	assert.Equal(t, sum, 67)

	sum = CountCheats(track, 64)
	assert.Equal(t, sum, 86)

	sum = CountCheats(track, 62)
	assert.Equal(t, sum, 106)

	sum = CountCheats(track, 60)
	assert.Equal(t, sum, 129)

	sum = CountCheats(track, 58)
	assert.Equal(t, sum, 154)

	sum = CountCheats(track, 56)
	assert.Equal(t, sum, 193)

	sum = CountCheats(track, 54)
	assert.Equal(t, sum, 222)

	sum = CountCheats(track, 52)
	assert.Equal(t, sum, 253)

	sum = CountCheats(track, 50)
	assert.Equal(t, sum, 285)
}

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	track := ParseInput(content)

	sum := CountCheats(track, 100)
	assert.Equal(t, sum, 983054)
}
