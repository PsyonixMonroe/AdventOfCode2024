package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSamplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	if content == "" {
		t.Fatal("Failed to read content")
	}

	cw := ReadCrossword(content)

	count := FindXmas(cw)
	assert.Equal(t, count, 18)
}

func TestCountAllWays(t *testing.T) {
	chars := []rune{
		'S', '.', '.', 'S', '.', '.', 'S',
		'.', 'A', '.', 'A', '.', 'A', '.',
		'.', '.', 'M', 'M', 'M', '.', '.',
		'S', 'A', 'M', 'X', 'M', 'A', 'S',
		'.', '.', 'M', 'M', 'M', '.', '.',
		'.', 'A', '.', 'A', '.', 'A', '.',
		'S', '.', '.', 'S', '.', '.', 'S',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	counts := CountAllWays(cw, 3, 3)
	assert.Equal(t, counts, 8)
	counts = FindXmas(cw)
	assert.Equal(t, counts, 8)
}

func TestCountLeft(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'S', 'A', 'M', 'X', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckHorizontalLeft(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountOverflowLeft(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', 'S',
		'A', 'M', 'X', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckHorizontalLeft(cw, 3, 2)
	assert.Equal(t, found, false)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 0)
}

func TestCountOverflowRight(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', 'X', 'M', 'A',
		'S', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckHorizontalRight(cw, 3, 4)
	assert.Equal(t, found, false)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 0)
}

func TestCountOverflowUp(t *testing.T) {
	chars := []rune{
		'.', '.', '.', 'A', '.', '.', '.',
		'.', '.', '.', 'M', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'S', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckVerticalUp(cw, 2, 3)
	assert.Equal(t, found, false)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 0)
}

func TestCountOverflowDown(t *testing.T) {
	chars := []rune{
		'.', '.', '.', 'S', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', 'M', '.', '.', '.',
		'.', '.', '.', 'A', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckVerticalDown(cw, 4, 3)
	assert.Equal(t, found, false)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 0)
}

func TestCountUp(t *testing.T) {
	chars := []rune{
		'.', '.', '.', 'S', '.', '.', '.',
		'.', '.', '.', 'A', '.', '.', '.',
		'.', '.', '.', 'M', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckVerticalUp(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountRight(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'X', 'M', 'A', 'S',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckHorizontalRight(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountDown(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', 'M', '.', '.', '.',
		'.', '.', '.', 'A', '.', '.', '.',
		'.', '.', '.', 'S', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckVerticalDown(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountDiagUpRight(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', 'S',
		'.', '.', '.', '.', '.', 'A', '.',
		'.', '.', '.', '.', 'M', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckDiagonalUpRight(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountDiagDownRight(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', '.', 'M', '.', '.',
		'.', '.', '.', '.', '.', 'A', '.',
		'.', '.', '.', '.', '.', '.', 'S',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckDiagonalDownRight(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountDiagDownLeft(t *testing.T) {
	chars := []rune{
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', 'M', '.', '.', '.', '.',
		'.', 'A', '.', '.', '.', '.', '.',
		'S', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckDiagonalDownLeft(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestCountDiagUpLeft(t *testing.T) {
	chars := []rune{
		'S', '.', '.', '.', '.', '.', '.',
		'.', 'A', '.', '.', '.', '.', '.',
		'.', '.', 'M', '.', '.', '.', '.',
		'.', '.', '.', 'X', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
	}
	cw := CrossWord{chars: chars, rows: 7, columns: 7}
	found := CheckDiagonalUpLeft(cw, 3, 3)
	assert.Equal(t, found, true)
	counts := FindXmas(cw)
	assert.Equal(t, counts, 1)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	if content == "" {
		t.Fatal("Failed to read content")
	}

	cw := ReadCrossword(content)

	count := FindXmas(cw)
	assert.Equal(t, count, 2562) // this is the accepted output, not sure why we get 1 additional
	// assert.Equal(t, count, 2563)
}

func TestSamplePart2(t *testing.T) {
	content := ReadInput("test.txt")
	if content == "" {
		t.Fatal("Failed to read content")
	}

	cw := ReadCrossword(content)
	count := FindXmas2(cw)
	assert.Equal(t, count, 9)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	if content == "" {
		t.Fatal("Failed to read content")
	}

	cw := ReadCrossword(content)
	count := FindXmas2(cw)
	assert.Equal(t, count, 1902)
}
