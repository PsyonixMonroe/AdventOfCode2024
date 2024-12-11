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

func RunStoneSimMemo(inStones []int, iterations int) int {
	stones := make(map[int]int, 0)
	for _, stone := range inStones {
		numStones, found := stones[stone]
		if found {
			stones[stone] = numStones + 1
		} else {
			stones[stone] = 1
		}
	}

	for range iterations {
		newStones := make(map[int]int, 0)
		for stone, count := range stones {
			if stone == 0 {
				newCount, found := newStones[1]
				if found {
					newStones[1] = newCount + count
				} else {
					newStones[1] = count
				}
				continue
			}
			if evenDigits(stone) {
				// split stone
				num1, num2 := splitDigits(stone)
				num1Count, num1Found := newStones[num1]
				if num1Found {
					newStones[num1] = num1Count + count
				} else {
					newStones[num1] = count
				}
				num2Count, num2Found := newStones[num2]
				if num2Found {
					newStones[num2] = num2Count + count
				} else {
					newStones[num2] = count
				}
				continue
			}
			newStone := stone * 2024
			newCount, found := newStones[newStone]
			if found {
				newStones[newStone] = newCount + count
			} else {
				newStones[newStone] = count
			}

		}
		stones = newStones
	}

	return sumStones(stones)
}

func sumStones(stones map[int]int) int {
	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
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
