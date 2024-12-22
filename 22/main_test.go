package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	seeds := GetSeeds(content)
	sum := 0
	for _, seed := range seeds {
		sum += GetNumForSeedIterations(seed, 2000)
	}

	assert.Equal(t, sum, 37327623)
}

func TestFirstSeedIterations(t *testing.T) {
	seeds := []int{123, 15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432, 5908254}
	for i, seed := range seeds {
		if i == len(seeds)-1 {
			continue
		}
		result := GetNumForSeedIterations(seed, 1)
		assert.Equal(t, result, seeds[i+1])
	}
}

func TestEachSeedPart1(t *testing.T) {
	seedResults := make(map[int]int)
	seedResults[1] = 8685429
	seedResults[10] = 4700978
	seedResults[100] = 15273692
	seedResults[2024] = 8667524

	for seed, result := range seedResults {
		calc := GetNumForSeedIterations(seed, 2000)
		assert.Equal(t, calc, result)
	}
}

func TestPrune(t *testing.T) {
	val := Prune(100000000)
	assert.Equal(t, val, 16113920)
}

func TestMix(t *testing.T) {
	val := Mix(15, 42)
	assert.Equal(t, val, 37)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	seeds := GetSeeds(content)
	sum := 0
	for _, seed := range seeds {
		sum += GetNumForSeedIterations(seed, 2000)
	}

	assert.Equal(t, sum, 18317943467)
}
