package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func GetAllSeedIterations(seed int, iterations int) []int {
	secret := seed
	allSecrets := []int{secret}
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
		allSecrets = append(allSecrets, secret)
	}

	return allSecrets
}

func Mix(val int, current int) int {
	return val ^ current
}

func Prune(val int) int {
	// note 16777216 is 2^23
	return val % 16777216
}

type Price struct {
	Actual int
	Diff   int
}

func GetPriceAndDiff(secrets []int) []Price {
	prices := []Price{}
	for i, secret := range secrets {
		diff := 0
		price := secret % 10
		if i != 0 {
			diff = price - prices[i-1].Actual
		}
		prices = append(prices, Price{Actual: price, Diff: diff})
	}

	return prices
}

func FindBestPatternSum(allPrices [][]Price) int {
	workerCtx, workerCancel := context.WithCancel(context.Background())
	priceChan := make(chan []int, 100)
	sumChan := make(chan int, 100)
	workerwg := sync.WaitGroup{}
	workerwg.Add(12)

	for range 12 {
		go func() {
			defer workerwg.Done()
			done := false
			for {
				select {
				case <-workerCtx.Done():
					done = true
				case pattern := <-priceChan:
					sum := SumAllPrices(pattern, allPrices)
					// if sum > 20 {
					// 	fmt.Fprintf(os.Stderr, "Found Good Sum with pattern: %d = (%d, %d, %d, %d)\n", sum, pattern[0], pattern[1], pattern[2], pattern[3])
					// }
					sumChan <- sum
				}
				if done {
					break
				}
			}
		}()
	}

	endCtx, endCancel := context.WithCancel(context.Background())
	endwg := sync.WaitGroup{}
	endwg.Add(1)
	best := 0
	go func() {
		defer endwg.Done()
		done := false
		for {
			select {
			case <-endCtx.Done():
				done = true
			case test := <-sumChan:
				if test > best {
					best = test
				}
			}
			if done {
				break
			}
		}
	}()

	// generate work
	for i := range 19 {
		a := i - 9
		for j := range 19 {
			b := j - 9
			for k := range 19 {
				c := k - 9
				for l := range 19 {
					d := l - 9
					priceChan <- []int{a, b, c, d}
				}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Finished Generating Work\n")
	// wait for workers to finish all processing
	for len(priceChan) > 0 {
	}
	workerCancel()
	workerwg.Wait()

	fmt.Fprintf(os.Stderr, "Worker Queue Empty\n")
	// wait for reducer to finish all processing
	for len(sumChan) > 0 {
	}
	endCancel()
	endwg.Wait()

	fmt.Fprintf(os.Stderr, "Reducer Queue Empty\n")
	// found the best
	return best
}

func SumAllPrices(pattern []int, allPrices [][]Price) int {
	sum := 0

	for _, merchant := range allPrices {
		for i := range len(merchant) - 4 {
			if merchant[i].Diff != pattern[0] {
				continue
			}
			if merchant[i+1].Diff != pattern[1] {
				continue
			}
			if merchant[i+2].Diff != pattern[2] {
				continue
			}
			if merchant[i+3].Diff != pattern[3] {
				continue
			}
			sum += merchant[i+3].Actual
			break
		}
	}

	return sum
}
