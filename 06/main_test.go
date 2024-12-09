package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSamplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	room := ParseInput(content)
	WalkMap(&room)
	count := CountPath(room)
	assert.Equal(t, count, 41)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	room := ParseInput(content)
	WalkMap(&room)
	count := CountPath(room)
	assert.Equal(t, count, 5516)
}

func TestSamplePart2(t *testing.T) {
	content := ReadInput("test.txt")
	room := ParseInputComplex(content)
	count := WalkAllMaps(room)
	assert.Equal(t, count, 6)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	room := ParseInputComplex(content)
	count := WalkAllMaps(room)
	assert.Equal(t, count, 2008)
}
