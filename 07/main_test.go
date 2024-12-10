package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	eq := ParseInput(content)
	sum := SumGoodEquations(eq)
	assert.Equal(t, sum, 3749)
}

func TestSumGoodEquations(t *testing.T) {
	equations := []Equation{}
	eq := Equation{solution: 190, args: []int{19, 10}}
	equations = append(equations, eq)
	sum := SumGoodEquations(equations)
	assert.Equal(t, sum, 190)
}

func TestSumGoodEquationsReverse(t *testing.T) {
	equations := []Equation{}
	eq := Equation{solution: 190, args: []int{10, 19}}
	equations = append(equations, eq)
	sum := SumGoodEquations(equations)
	assert.Equal(t, sum, 190)
}

func TestSlicing(t *testing.T) {
	value := ProcessEquation(190, 19, []int{10})
	assert.Equal(t, 190, value)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	eq := ParseInput(content)
	sum := SumGoodEquations(eq)
	assert.Equal(t, sum, 2654749936343)
}

func TestSimplePart2(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	eq := ParseInput(content)
	sum := SumGoodEquations2(eq)
	assert.Equal(t, sum, 11387)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	eq := ParseInput(content)
	sum := SumGoodEquations2(eq)
	assert.Equal(t, sum, 124060392153684)
}

func TestConcat(t *testing.T) {
	value := concat(8, 6)
	assert.Equal(t, value, 86)
}
