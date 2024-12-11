package main

import (
	"fmt"
	"os"
	"strconv"
)

func RunStonesSim(stones []int, iterations int) []int {
	for range iterations {
		newStones := []int{}
		for i, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}
			if evenDigits(stone) {
				// split stone
				num1, num2 := splitDigits(stone)
				newStones = append(newStones, num1)
				newStones = append(newStones, num2)
				continue
			}
			newStones = append(newStones, stone*2024)
			stones[i] *= 2024
		}
		stones = newStones
	}
	return stones
}

func evenDigits(num int) bool {
	numStr := strconv.Itoa(num)
	return len(numStr)%2 == 0
}

func splitDigits(num int) (int, int) {
	numStr := strconv.Itoa(num)
	num1, err := strconv.Atoi(numStr[:len(numStr)/2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse num1 from %s, len: %d\n", numStr, len(numStr)/2)
		return -1, -1
	}

	num2, err := strconv.Atoi(numStr[len(numStr)/2:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse num2 from %s, len: %d\n", numStr, len(numStr)/2)
		return -1, -1
	}

	return num1, num2
}
