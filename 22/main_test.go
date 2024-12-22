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

func TestFirstSeedIterationsPart1(t *testing.T) {
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

func TestFirstSeedIterationsPart2(t *testing.T) {
	seeds := []int{123, 15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432}
	for i, seed := range seeds {
		if i == len(seeds)-1 {
			continue
		}
		result := GetNumForSeedIterations(seed, 1)
		assert.Equal(t, result, seeds[i+1])
	}

	prices := GetPriceAndDiff(seeds)
	expectedPrices := []Price{{Actual: 3, Diff: 0}, {Actual: 0, Diff: -3}, {Actual: 6, Diff: 6}, {Actual: 5, Diff: -1}, {Actual: 4, Diff: -1}, {Actual: 4, Diff: 0}, {Actual: 6, Diff: 2}, {Actual: 4, Diff: -2}, {Actual: 4, Diff: 0}, {Actual: 2, Diff: -2}}
	assert.Equal(t, len(prices), len(expectedPrices))
	for i, price := range prices {
		assert.Equal(t, price.Actual, expectedPrices[i].Actual)
		assert.Equal(t, price.Diff, expectedPrices[i].Diff)
	}
}

func TestGetAllSeedIterations(t *testing.T) {
	expected := []int{123, 15887950, 16495136, 527345, 704524, 1553684, 12683156, 11100544, 12249484, 7753432}
	seeds := GetAllSeedIterations(123, 9)
	assert.Equal(t, len(seeds), len(expected))
	for i := range seeds {
		assert.Equal(t, seeds[i], expected[i])
	}
}

func TestGenerator(t *testing.T) {
	expected := []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := []int{}
	for i := range 19 {
		actual = append(actual, i-9)
	}

	assert.Equal(t, len(actual), len(expected))
	for i := range actual {
		assert.Equal(t, actual[i], expected[i])
	}
}

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	seeds := GetSeeds(content)
	prices := [][]Price{}
	for _, seed := range seeds {
		secrets := GetAllSeedIterations(seed, 2000)
		merchantPrices := GetPriceAndDiff(secrets)
		prices = append(prices, merchantPrices)
	}
	sum := FindBestPatternSum(prices)

	assert.Equal(t, sum, 23)
}

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	seeds := GetSeeds(content)
	prices := [][]Price{}
	for _, seed := range seeds {
		secrets := GetAllSeedIterations(seed, 2000)
		merchantPrices := GetPriceAndDiff(secrets)
		prices = append(prices, merchantPrices)
	}
	sum := FindBestPatternSum(prices)

	assert.Equal(t, sum, 2018)
}
