package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	m := ParseInput(content)
	MarkPoints(&m)
	sum := CountLocations(m)
	assert.Equal(t, sum, 14)
}

func TestIntersectionFromStation(t *testing.T) {
	x1, y1, x2, y2 := IntersectionsFromStation(2, 2, 3, 3)
	assert.Equal(t, x1, 1)
	assert.Equal(t, y1, 1)
	assert.Equal(t, x2, 4)
	assert.Equal(t, y2, 4)
}

func TestIntersectionFromStationRev(t *testing.T) {
	x1, y1, x2, y2 := IntersectionsFromStation(2, 3, 3, 2)
	assert.Equal(t, x1, 1)
	assert.Equal(t, y1, 4)
	assert.Equal(t, x2, 4)
	assert.Equal(t, y2, 1)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	m := ParseInput(content)
	MarkPoints(&m)
	sum := CountLocations(m)
	assert.Equal(t, sum, 361)
}
