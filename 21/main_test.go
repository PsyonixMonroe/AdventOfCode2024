package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	input1 := []rune{'0', '2', '9', 'A'}
	keypress1 := GetInputForCode(input1)
	complexity1 := GetComplexityForCode(input1, int64(len(keypress1)))
	assert.Equal(t, complexity1, int64(1972))

	input2 := []rune{'9', '8', '0', 'A'}
	keypress2 := GetInputForCode(input2)
	complexity2 := GetComplexityForCode(input2, int64(len(keypress2)))
	assert.Equal(t, complexity2, int64(58800))

	input3 := []rune{'1', '7', '9', 'A'}
	keypress3 := GetInputForCode(input3)
	complexity3 := GetComplexityForCode(input3, int64(len(keypress3)))
	assert.Equal(t, complexity3, int64(12172))

	input4 := []rune{'4', '5', '6', 'A'}
	keypress4 := GetInputForCode(input4)
	complexity4 := GetComplexityForCode(input4, int64(len(keypress4)))
	assert.Equal(t, complexity4, int64(29184))

	input5 := []rune{'3', '7', '9', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, int64(len(keypress5)))
	assert.Equal(t, complexity5, int64(24256))

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, int64(126384))
}

func TestSimpleSinglePart1(t *testing.T) {
	input := []rune{'0', '2', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	assert.Equal(t, complexity, int64(1972))
}

func TestSimplePart1Num5(t *testing.T) {
	input5 := []rune{'3', '7', '9', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, int64(len(keypress5)))
	assert.Equal(t, complexity5, int64(24256))
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
	complexity1 := GetComplexityForCode(input1, int64(len(keypress1)))

	input2 := []rune{'9', '8', '5', 'A'}
	keypress2 := GetInputForCode(input2)
	complexity2 := GetComplexityForCode(input2, int64(len(keypress2)))

	input3 := []rune{'3', '4', '0', 'A'}
	keypress3 := GetInputForCode(input3)
	complexity3 := GetComplexityForCode(input3, int64(len(keypress3)))

	input4 := []rune{'4', '8', '9', 'A'}
	keypress4 := GetInputForCode(input4)
	complexity4 := GetComplexityForCode(input4, int64(len(keypress4)))

	input5 := []rune{'9', '6', '4', 'A'}
	keypress5 := GetInputForCode(input5)
	complexity5 := GetComplexityForCode(input5, int64(len(keypress5)))

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, int64(215374))
}

func TestFullPart1Num1(t *testing.T) {
	input := []rune{'3', '1', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	assert.Equal(t, complexity, int64(22330))
}

func TestFullPart1Num2(t *testing.T) {
	input := []rune{'9', '8', '5', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	assert.Equal(t, complexity, int64(65010))
}

func TestFullPart1Num3(t *testing.T) {
	input := []rune{'3', '4', '0', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	assert.Equal(t, complexity, int64(22440))
}

func TestFullPart1Num4(t *testing.T) {
	input := []rune{'4', '8', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	assert.Equal(t, complexity, int64(36186))
}

func TestSimplePart1_redux(t *testing.T) {
	input1 := []rune{'0', '2', '9', 'A'}
	keypress1 := GetBestNumeric(input1)
	len1 := RunIterations(keypress1, 2)
	complexity1 := GetComplexityForCode(input1, len1)
	assert.Equal(t, complexity1, int64(1972))

	input2 := []rune{'9', '8', '0', 'A'}
	keypress2 := GetBestNumeric(input2)
	len2 := RunIterations(keypress2, 2)
	complexity2 := GetComplexityForCode(input2, len2)
	assert.Equal(t, complexity2, int64(58800))

	input3 := []rune{'1', '7', '9', 'A'}
	keypress3 := GetBestNumeric(input3)
	len3 := RunIterations(keypress3, 2)
	complexity3 := GetComplexityForCode(input3, len3)
	assert.Equal(t, complexity3, int64(12172))

	input4 := []rune{'4', '5', '6', 'A'}
	keypress4 := GetBestNumeric(input4)
	len4 := RunIterations(keypress4, 2)
	complexity4 := GetComplexityForCode(input4, len4)
	assert.Equal(t, complexity4, int64(29184))

	input5 := []rune{'3', '7', '9', 'A'}
	keypress5 := GetBestNumeric(input5)
	len5 := RunIterations(keypress5, 2)
	complexity5 := GetComplexityForCode(input5, len5)
	assert.Equal(t, complexity5, int64(24256))

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, int64(126384))
}

func TestFullPart1Num5(t *testing.T) {
	input := []rune{'9', '6', '4', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))

	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)
	assert.Equal(t, complexity, int64(69408))
}

func TestFullPart1Num1P2Solve(t *testing.T) {
	input := []rune{'3', '1', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)
	assert.Equal(t, complexity, int64(22330))
}

func TestFullPart1Num2P2Solve(t *testing.T) {
	input := []rune{'9', '8', '5', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)

	assert.Equal(t, complexity, int64(65010))
}

func TestFullPart1Num3P2Solve(t *testing.T) {
	input := []rune{'3', '4', '0', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)
	assert.Equal(t, complexity, int64(22440))
}

func TestFullPart1Num4P2Solve(t *testing.T) {
	input := []rune{'4', '8', '9', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)
	assert.Equal(t, complexity, int64(36186))
}

func TestFullPart1Num5P2Solve(t *testing.T) {
	input := []rune{'9', '6', '4', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))

	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)
	assert.Equal(t, complexity, int64(69408))
}

func TestSimpleSinglePart2(t *testing.T) {
	input := []rune{'0', '2', '9', 'A'}
	keypress := GetBestNumeric(input)
	numPresses := RunIterations(keypress, 2)
	complexity := GetComplexityForCode(input, numPresses)
	assert.Equal(t, complexity, int64(1972))
}

func TestInput3Part2(t *testing.T) {
	input3 := []rune{'1', '7', '9', 'A'}
	keypress3 := GetInputForCode(input3)
	complexity3 := GetComplexityForCode(input3, int64(len(keypress3)))
	assert.Equal(t, complexity3, int64(12172))

	keypress3 = GetBestNumeric(input3)
	len3 := RunIterations(keypress3, 2)
	complexity3 = GetComplexityForCode(input3, len3)

	assert.Equal(t, complexity3, int64(12172))
}

func TestFullPart2(t *testing.T) {
	var complexity5 int64
	var complexity4 int64
	var complexity3 int64
	var complexity2 int64
	var complexity1 int64
	input1 := []rune{'3', '1', '9', 'A'}
	keypress1 := GetInputForCode25(input1)
	complexity1 = GetComplexityForCode(input1, keypress1)

	input2 := []rune{'9', '8', '5', 'A'}
	keypress2 := GetInputForCode25(input2)
	complexity2 = GetComplexityForCode(input2, keypress2)

	input3 := []rune{'3', '4', '0', 'A'}
	keypress3 := GetInputForCode25(input3)
	complexity3 = GetComplexityForCode(input3, keypress3)

	input4 := []rune{'4', '8', '9', 'A'}
	keypress4 := GetInputForCode25(input4)
	complexity4 = GetComplexityForCode(input4, keypress4)

	input5 := []rune{'9', '6', '4', 'A'}
	keypress5 := GetInputForCode25(input5)
	complexity5 = GetComplexityForCode(input5, keypress5)

	sum := complexity1 + complexity2 + complexity3 + complexity4 + complexity5
	assert.Equal(t, sum, int64(260586897262600))
}

func TestFullPart1Num2P2Solve_redux(t *testing.T) {
	input := []rune{'9', '8', '5', 'A'}
	keypress := GetInputForCode(input)
	complexity := GetComplexityForCode(input, int64(len(keypress)))
	newInput := GetBestNumeric(input)
	newLen := RunIterations(newInput, 2)
	newComplex := GetComplexityForCode(input, newLen)
	assert.Equal(t, complexity, newComplex)

	assert.Equal(t, complexity, int64(65010))
}
