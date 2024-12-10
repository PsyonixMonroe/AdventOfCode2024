package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 1)
}

func TestSimple2Part1(t *testing.T) {
	content := ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 2)
}

func TestSimple3Part1(t *testing.T) {
	content := ReadInput("test3.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 4)
}

func TestSimple4Part1(t *testing.T) {
	content := ReadInput("test4.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 3)
}

func TestSimple5Part1(t *testing.T) {
	content := ReadInput("test5.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 36)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindTrails(m)
	assert.Equal(t, sum, 430)
}

func TestSample5Part2(t *testing.T) {
	content := ReadInput("test5.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindScoreTrailHeads(m)
	assert.Equal(t, sum, 81)
}

func TestSample6Part2(t *testing.T) {
	content := ReadInput("test6.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindScoreTrailHeads(m)
	assert.Equal(t, sum, 3)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	m := ParseTopoMap(content)
	sum := FindScoreTrailHeads(m)
	assert.Equal(t, sum, 928)
}
