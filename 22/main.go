package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetSeeds(content string) []int {
	lines := strings.Split(strings.Trim(content, "\n \t"), "\n")
	values := []int{}
	for _, line := range lines {
		strVal := strings.Trim(line, "\n \t")
		val, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse input line: %s\n", strVal)
		}

		values = append(values, val)
	}
	return values
}

func GetNumForSeedIterations(seed int, iterations int) int {
	secret := seed
	for range iterations {
		next := secret * 64
		secret = Mix(next, secret)
		secret = Prune(secret)
		next = secret / 32
		secret = Mix(next, secret)
		secret = Prune(secret)
		next = secret * 2048
		secret = Mix(next, secret)
		secret = Prune(secret)
	}

	return secret
}

func Mix(val int, current int) int {
	return val ^ current
}

func Prune(val int) int {
	// note 16777216 is 2^23
	return val % 16777216
}
