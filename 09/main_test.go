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

func TestGetFreeSpace(t *testing.T) {
	disk := []int{2, 2, 2, -1, -1, 3, -1, -1, -1, 4, 5}
	freeSpaceStart := GetFreeSpace(disk, 3, 8)
	assert.Equal(t, freeSpaceStart, 6)
}

func TestSwap(t *testing.T) {
	disk := []int{2, 2, 2, -1, -1, 3, -1, -1, -1, 4, 4, 5}
	swapFile(&disk, 9, 10, 3)
	afterSwap := []int{2, 2, 2, 4, 4, 3, -1, -1, -1, -1, -1, 5}
	assertEqual(t, disk, afterSwap)
}

func TestGetNextFileRange(t *testing.T) {
	disk := []int{2, 2, 2, -1, -1, 3, -1, -1, -1, 4, 4, 5}
	start4, end4 := GetNextFileRange(disk, 10)
	assert.Equal(t, start4, 9)
	assert.Equal(t, end4, 10)

	start3, end3 := GetNextFileRange(disk, 8)
	assert.Equal(t, start3, 5)
	assert.Equal(t, end3, 5)

	start2, end2 := GetNextFileRange(disk, 4)
	assert.Equal(t, start2, 0)
	assert.Equal(t, end2, 2)

	start9, end9 := GetNextFileRange(disk, 11)
	assert.Equal(t, start9, 11)
	assert.Equal(t, end9, 11)
}

func TestSimplePart2(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	disk := ParseInput(content)
	DeFragFilesContiguous(&disk)

	checkSum := Checksum(disk)
	assert.Equal(t, checkSum, 2858)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	disk := ParseInput(content)
	DeFragFilesContiguous(&disk)

	checkSum := Checksum(disk)
	assert.Equal(t, checkSum, 6349492251099)
}
