package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimple3Part1(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score := FindBestPath(course)
	assert.Equal(t, score, 2008)
}

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score := FindBestPath(course)
	assert.Equal(t, score, 7036)
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score := FindBestPath(course)
	assert.Equal(t, score, 11048)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score := FindBestPath(course)
	assert.Equal(t, score, 98416)
}
