package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func swap(disk *[]int, i int, j int) {
	tmp := (*disk)[i]
	(*disk)[i] = (*disk)[j]
	(*disk)[j] = tmp
}

func ParseInput(content string) []int {
	disk := make([]int, 0)
	for pos, size := range strings.Trim(content, "\n") {
		num, err := strconv.Atoi(string(size))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid disk size: %s\n", string(size))
			continue
		}

		var item int
		if pos%2 == 0 {
			// file
			item = pos / 2
		} else {
			// empty space
			item = -1
		}
		for range num {
			disk = append(disk, item)
		}
	}

	return disk
}

func DeFragFiles(disk *[]int) {
	freeSpaceLoc := GetNextFreeSpaceLoc(*disk, 0)
	for i := range len(*disk) {
		moveFileLoc := len(*disk) - 1 - i
		if moveFileLoc <= freeSpaceLoc {
			// done
			break
		}
		if (*disk)[moveFileLoc] == -1 {
			continue
		}
		swap(disk, moveFileLoc, freeSpaceLoc)
		freeSpaceLoc = GetNextFreeSpaceLoc(*disk, freeSpaceLoc+1)
	}
}

func GetNextFreeSpaceLoc(disk []int, pos int) int {
	maxLen := len(disk)
	for i := range maxLen {
		newPos := pos + i
		if newPos >= maxLen {
			return maxLen
		}
		if disk[newPos] == -1 {
			// found free space
			return newPos
		}
	}
	return maxLen
}

func Checksum(disk []int) int {
	sum := 0
	for i, fileID := range disk {
		if fileID < 0 {
			continue
		}
		sum += i * fileID
	}
	return sum
}
