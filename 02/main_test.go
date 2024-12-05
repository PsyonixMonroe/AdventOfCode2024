package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsSafeAscSimple(t *testing.T) {
	report := Report{num: 1, levels: []int{1, 3, 5, 7, 9}}
	test := IsSafeAsc(report)
	assert.Equal(t, test, true)
}

func TestIsSafeAscFailDesc(t *testing.T) {
	report := Report{num: 1, levels: []int{9, 7, 5, 3, 1}}
	test := IsSafeAsc(report)
	assert.Equal(t, test, false)
}

func TestIsSafeAscFailRange(t *testing.T) {
	report := Report{num: 1, levels: []int{1, 3, 5, 7, 12}}
	test := IsSafeAsc(report)
	assert.Equal(t, test, false)
}

func TestIsSafeAscFailRand(t *testing.T) {
	report := Report{num: 1, levels: []int{1, 3, 15, 7, 9}}
	test := IsSafeAsc(report)
	assert.Equal(t, test, false)
}

func TestIsSafeDescSimple(t *testing.T) {
	report := Report{num: 1, levels: []int{9, 7, 5, 3, 1}}
	test := IsSafeDesc(report)
	assert.Equal(t, test, true)
}

func TestIsSafeDescFailAsc(t *testing.T) {
	report := Report{num: 1, levels: []int{1, 3, 5, 7, 9}}
	test := IsSafeDesc(report)
	assert.Equal(t, test, false)
}

func TestIsSafeDescFailRange(t *testing.T) {
	report := Report{num: 1, levels: []int{15, 7, 5, 3, 1}}
	test := IsSafeDesc(report)
	assert.Equal(t, test, false)
}

func TestIsSafeDescFailRand(t *testing.T) {
	report := Report{num: 1, levels: []int{8, 7, 15, 3, 1}}
	test := IsSafeDesc(report)
	assert.Equal(t, test, false)
}

func TestPart1Sample(t *testing.T) {
	reports := ParseInput("test.txt")
	if len(reports) == 0 {
		t.Fatal("Failed to read in input file")
	}

	count := CountSafe(reports)
	assert.Equal(t, count, 2)
}

func TestPart1Full(t *testing.T) {
	reports := ParseInput("input.txt")
	if len(reports) == 0 {
		t.Fatal("Failed to read in input file")
	}

	count := CountSafe(reports)
	assert.Equal(t, count, 257)
}

func TestPart2Sample(t *testing.T) {
	reports := ParseInput("test.txt")
	if len(reports) == 0 {
		t.Fatal("Failed to read in input file")
	}

	count := CountSafeDamp(reports)
	assert.Equal(t, count, 4)
}

func TestPart2Full(t *testing.T) {
	reports := ParseInput("input.txt")
	if len(reports) == 0 {
		t.Fatal("Failed to read in input file")
	}

	count := CountSafeDamp(reports)
	assert.Equal(t, count, 257)
}
