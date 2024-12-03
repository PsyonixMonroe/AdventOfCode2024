package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World")
}

func ReadInput(filename string) ([]int, []int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Read file: %s\n", filename)
		return []int{}, []int{}
	}

	left := []int{}
	right := []int{}
	input := string(content)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		splits := strings.Fields(strings.Trim(line, "\n"))
		if len(splits) != 2 {
			fmt.Fprintf(os.Stderr, "Improperly formatted input: '%s'\n", strings.Trim(line, "\n"))
			continue
		}
		leftNum, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse leftNum for line: %s\n", splits[0])
		}
		left = append(left, leftNum)
		rightNum, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse rightNum for line: %s\n", splits[1])
		}
		right = append(right, rightNum)
	}

	return left, right
}

func GetDiff(left []int, right []int) int {
	diff := 0
	for i, leftNum := range left {
		rightNum := right[i]
		diff += intAbs(leftNum - rightNum)
	}

	return diff
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
