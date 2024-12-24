package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(2024))
}

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(4))
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(43559017878162))
}
