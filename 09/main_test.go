package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	disk := ParseInput(content)
	DeFragFiles(&disk)

	checkSum := Checksum(disk)
	assert.Equal(t, checkSum, 1928)
}

func TestParseInput(t *testing.T) {
	disk := ParseInput("12345")
	expectedDisk := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	assertEqual(t, disk, expectedDisk)
}

func assertEqual(t *testing.T, actual []int, expected []int) {
	assert.Equal(t, len(actual), len(expected))
	for i, val := range actual {
		assert.Equal(t, val, expected[i])
	}
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	disk := ParseInput(content)
	DeFragFiles(&disk)

	checkSum := Checksum(disk)
	assert.Equal(t, checkSum, 6334655979668)
}

