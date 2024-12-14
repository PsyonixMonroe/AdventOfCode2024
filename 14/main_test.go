package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	drones := ParseDrones(content)
	assert.NotEqual(t, len(drones), 0)

	drones = SimulateDrones(drones, 7, 11, 100)
	assert.NotEqual(t, len(drones), 0)

	score := CalculateSafety(drones, 7, 11)
	assert.Equal(t, score, 12)
}

func TestSingleDronePart1(t *testing.T) {
	drone := Drone{locX: 4, locY: 2, velX: -3, velY: 2}
	drones := []Drone{drone}

	drones = SimulateDrones(drones, 7, 11, 5)
	assert.Equal(t, len(drones), 1)
	assert.Equal(t, drones[0].locX, 3)
	assert.Equal(t, drones[0].locY, 1)
	assert.Equal(t, drones[0].velX, -3)
	assert.Equal(t, drones[0].velY, 2)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	drones := ParseDrones(content)

	drones = SimulateDrones(drones, 103, 101, 100)

	score := CalculateSafety(drones, 103, 101)
	assert.Equal(t, score, 218433348)
}

func TestFindTreePart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	drones := ParseDrones(content)

	_ = SimulateDrones(drones, 103, 101, 6511)
}
