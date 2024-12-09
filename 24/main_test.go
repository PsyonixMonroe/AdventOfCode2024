package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
}
