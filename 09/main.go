package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
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

func DeFragFilesContiguous(disk *[]int) {
	i := len(*disk) - 1
	for {
		fileStart, fileEnd := GetNextFileRange(*disk, i)
		if fileStart < 0 {
			break
		}
		freeSpaceStart := GetFreeSpace(*disk, fileEnd-fileStart+1, fileStart)
		fmt.Fprintf(os.Stderr, "")
		if freeSpaceStart > 0 {
			swapFile(disk, fileStart, fileEnd, freeSpaceStart)
		}
		i = fileStart - 1
		if i < 0 {
			break
		}
	}
}

func swapFile(disk *[]int, start int, end int, freeSpaceStart int) {
	count := 0
	for _, i := range lib.GetRange(start, end+1) {
		swap(disk, i, freeSpaceStart+count)
		count++
	}
}

func GetNextFileRange(disk []int, startPos int) (int, int) {
	for _, endVal := range lib.GetRevRange(0, startPos+1) {
		if disk[endVal] == -1 {
			continue
		}
		fileID := disk[endVal]
		for _, startVal := range lib.GetRevRange(0, endVal) {
			if disk[startVal] == -1 || fileID != disk[startVal] {
				return startVal + 1, endVal
			}
		}
		return 0, endVal
	}
	fmt.Fprintf(os.Stderr, "Didn't find file from %d\n", startPos)
	return -1, -1
}

func GetFreeSpace(disk []int, size int, fileLoc int) int {
	for i := range len(disk) {
		if i >= fileLoc {
			return -1
		}
		if disk[i] == -1 {
			found := true
			for j := range size {
				if disk[i+j] != -1 {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
