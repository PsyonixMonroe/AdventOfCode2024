package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	num    int
	levels []int
}

func ParseInput(filename string) []Report {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse input file %v\n", err)
		return []Report{}
	}

	reports := []Report{}
	for i, line := range strings.Split(strings.Trim(string(content), "\n"), "\n") {
		if line == "" {
			continue
		}
		valuesStr := strings.Fields(strings.Trim(line, "\n"))
		levels := []int{}
		for _, v := range valuesStr {
			level, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[%d] Unable to parse level, bad format: %s\n", i, line)
				continue
			}

			levels = append(levels, level)
		}

		reports = append(reports, Report{num: i, levels: levels})
	}

	return reports
}

func IsAsc(report Report) bool {
	for i, level := range report.levels {
		if i == 0 {
			continue
		}
		prevLevel := report.levels[i-1]
		if prevLevel == level {
			continue
		}
		return prevLevel < level
	}

	fmt.Fprintf(os.Stderr, "Unable to determine asc\n")
	return true
}

func CountSafe(reports []Report) int {
	safeCount := 0
	for _, report := range reports {
		isAsc := IsAsc(report)
		var isSafe bool
		if isAsc {
			isSafe = IsSafeAsc(report)
		} else {
			isSafe = IsSafeDesc(report)
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func IsSafeAsc(report Report) bool {
	for i, level := range report.levels {
		if i == 0 {
			continue
		}
		prevLevel := report.levels[i-1]
		diff := level - prevLevel
		// asc so diff should be positive between 1 and 3
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func IsSafeDesc(report Report) bool {
	for i, level := range report.levels {
		if i == 0 {
			continue
		}
		prevLevel := report.levels[i-1]
		diff := level - prevLevel
		// desc so diff should be negative between -1 and -3
		if diff < -3 || diff > -1 {
			return false
		}
	}

	return true
}
