package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSamplePart1(t *testing.T) {
	content := ReadInput("test.txt")

	rules, updates := ParseInput(content)

	count := CountGoodUpdates(rules, updates)

	assert.Equal(t, count, 143)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")

	rules, updates := ParseInput(content)

	count := CountGoodUpdates(rules, updates)

	assert.Equal(t, count, 4609)
}

func TestSamplePart2(t *testing.T) {
	content := ReadInput("test.txt")

	rules, updates := ParseInput(content)

	count := FixBadUpdatesAndCount(rules, updates)

	assert.Equal(t, count, 123)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")

	rules, updates := ParseInput(content)

	count := FixBadUpdatesAndCount(rules, updates)

	assert.Equal(t, count, 5723)
}
