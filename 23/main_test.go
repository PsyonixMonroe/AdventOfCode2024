package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	edges := ParseConnections(content)
	index := IndexFromEdges(edges)
	sum := CountNetworkedComputers(index)

	assert.Equal(t, sum, 7)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	edges := ParseConnections(content)
	index := IndexFromEdges(edges)
	sum := CountNetworkedComputers(index)

	assert.Equal(t, sum, 1064)
}
