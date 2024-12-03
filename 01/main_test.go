package main

import (
	"slices"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSample1(t *testing.T) {
	t.Logf("\n==== Testing Sample input for Problem 1 ====\n")

	left, right := ReadInput("test.txt")
	if len(left) == 0 {
		t.Fatalf("Didn't get any left input\n")
	}
	if len(right) == 0 {
		t.Fatalf("Didn't get any right input\n")
	}
	if len(left) != len(right) {
		t.Fatalf("Left and Right have different number of items: %d - %d\n", len(left), len(right))
	}
	left = slices.Sorted[int](slices.Values(left))
	right = slices.Sorted[int](slices.Values(right))
	diff := GetDiff(left, right)
	assert.Equal(t, diff, 11)
}

func TestFull1(t *testing.T) {
	t.Logf("\n==== Testing Sample input for Problem 1 ====\n")

	left, right := ReadInput("input.txt")
	if len(left) == 0 {
		t.Fatalf("Didn't get any left input\n")
	}
	if len(right) == 0 {
		t.Fatalf("Didn't get any right input\n")
	}
	if len(left) != len(right) {
		t.Fatalf("Left and Right have different number of items: %d - %d\n", len(left), len(right))
	}
	left = slices.Sorted[int](slices.Values(left))
	right = slices.Sorted[int](slices.Values(right))
	diff := GetDiff(left, right)
	assert.Equal(t, diff, 2367773)
}
