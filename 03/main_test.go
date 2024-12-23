package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSamplePart1(t *testing.T) {
	input := ReadInput("test.txt")
	commands := ParseMult(input)
	sum := ProcessCommands(commands)
	assert.Equal(t, sum, 161)
}

func TestFullPart1(t *testing.T) {
	input := ReadInput("input.txt")
	commands := ParseMult(input)
	sum := ProcessCommands(commands)
	assert.Equal(t, sum, 183788984)
}

func TestSamplePart2(t *testing.T) {
	input := ReadInput("test2.txt")
	commands := ParseMult2(input)
	sum := ProcessCommands(commands)
	assert.Equal(t, sum, 48)
}

func TestFullPart2(t *testing.T) {
	input := ReadInput("input.txt")
	commands := ParseMult2(input)
	sum := ProcessCommands(commands)
	assert.Equal(t, sum, 62098619)
}
