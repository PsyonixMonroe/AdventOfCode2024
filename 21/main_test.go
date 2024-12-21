package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	input1 := []rune{'0', '2', '9', 'A'}
	keypress1 := GetInputForCode(input1)
	complexity1 := GetComplexityForCode(input1, keypress1)
	assert.Equal(t, complexity1, 1972)

	input2 := []rune{'9', '8', '0', 'A'}
	keypress2 := GetInputForCode(input2)
	complexity2 := GetComplexityForCode(input2, keypress2)
	assert.Equal(t, complexity2, 58800)

	input3 := []rune{'1', '7', '9', 'A'}
	keypress3 := GetInputForCode(input3)
	complexity3 := GetComplexityForCode(input3, keypress3)
	assert.Equal(t, complexity3, 12172)

	input4 := []rune{'4', '5', '6', 'A'}
	keypress4 := GetInputForCode(input4)
	complexity4 := GetComplexityForCode(input4, keypress4)
	assert.Equal(t, complexity4, 29184)

	input5 := []rune{'3', '7', '9', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, keypress5)
	assert.Equal(t, complexity5, 24256)

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, 126384)
}

func TestSimpleSinglePart1(t *testing.T) {
	input := []rune{'0', '2', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 1972)
}

func TestSimplePart1Num5(t *testing.T) {
	input5 := []rune{'3', '7', '9', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, keypress5)
	assert.Equal(t, complexity5, 24256)
}

func TestSingleInputPart1(t *testing.T) {
	input := []rune{'0'}
	first := TranslateFromNumeric(input)
	assertAllEqual(t, first, []rune{'<', 'A'})
}

func assertAllEqual(t *testing.T, actual []rune, expected []rune) {
	assert.Equal(t, len(actual), len(expected))

	for i, c := range actual {
		assert.Equal(t, c, expected[i])
	}
}

func TestFullPart1(t *testing.T) {
	input1 := []rune{'3', '1', '9', 'A'}
	keypress1 := GetInputForCode(input1)
	complexity1 := GetComplexityForCode(input1, keypress1)

	input2 := []rune{'9', '8', '5', 'A'}
	keypress2 := GetInputForCode(input2)
	complexity2 := GetComplexityForCode(input2, keypress2)

	input3 := []rune{'3', '4', '0', 'A'}
	keypress3 := GetInputForCode(input3)
	complexity3 := GetComplexityForCode(input3, keypress3)

	input4 := []rune{'4', '8', '9', 'A'}
	keypress4 := GetInputForCode(input4)
	complexity4 := GetComplexityForCode(input4, keypress4)

	input5 := []rune{'9', '6', '4', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, keypress5)

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, 215374)
}

func TestFullPart1Num1(t *testing.T) {
	input := []rune{'3', '1', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 22330)
}

func TestFullPart1Num2(t *testing.T) {
	input := []rune{'9', '8', '5', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 65010)
}

func TestFullPart1Num3(t *testing.T) {
	input := []rune{'3', '4', '0', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 22440)
}

func TestFullPart1Num4(t *testing.T) {
	input := []rune{'4', '8', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 36186)
}

func TestFullPart1Num5(t *testing.T) {
	input := []rune{'9', '6', '4', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, keypress)
	assert.Equal(t, complexity, 69408)
}
