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

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	edges := ParseConnections(content)
	index := IndexFromEdges(edges)

	password := FindLanPassword(index)

	assert.Equal(t, password, "co,de,ka,ta")
}

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	edges := ParseConnections(content)
	index := IndexFromEdges(edges)

	password := FindLanPassword(index)

	assert.Equal(t, password, "aq,cc,ea,gc,jo,od,pa,rg,rv,ub,ul,vr,yy")
}
