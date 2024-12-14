package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Drone struct {
	locX int
	locY int
	velX int
	velY int
}

func (d Drone) String() string {
	return fmt.Sprintf("Drone{x: %d, y: %d, velX: %d, velY: %d}", d.locX, d.locY, d.velX, d.velY)
}

func ParseDrones(content string) []Drone {
	droneRegex := regexp.MustCompile(`p=([-\d]+),([-\d]+) v=([-\d]+),([-\d]+)`)
	drones := []Drone{}
	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		match := droneRegex.FindStringSubmatch(line)
		if len(match) != 5 {
			fmt.Fprintf(os.Stderr, "Unable to parse line: %s\n", line)
			continue
		}
		locY := lib.StringToInt(match[1])
		locX := lib.StringToInt(match[2])
		velY := lib.StringToInt(match[3])
		velX := lib.StringToInt(match[4])
		drones = append(drones, Drone{locX: locX, locY: locY, velX: velX, velY: velY})
	}
	return drones
}

func SimulateDrones(drones []Drone, height int, width int, iterations int) []Drone {
	for range iterations {
		newDrones := []Drone{}
		for _, drone := range drones {
			newLocX := drone.locX + drone.velX
			if newLocX < 0 {
				newLocX = height + newLocX
			}
			locX := newLocX % height

			newLocY := drone.locY + drone.velY
			if newLocY < 0 {
				newLocY = width + newLocY
			}
			locY := newLocY % width

			newDrones = append(newDrones, Drone{locX: locX, locY: locY, velX: drone.velX, velY: drone.velY})
		}
		drones = newDrones
	}
	return drones
}

func CalculateSafety(drones []Drone, height int, width int) int {
	midpointX := height / 2
	midpointY := width / 2
	q0 := 0
	q1 := 0
	q2 := 0
	q3 := 0

	for _, drone := range drones {
		if drone.locX == midpointX || drone.locY == midpointY {
			continue
		}
		if drone.locX < midpointX {
			if drone.locY < midpointY {
				q0++
			} else {
				q1++
			}
		} else {
			if drone.locY < midpointY {
				q2++
			} else {
				q3++
			}
		}
	}

	return q0 * q1 * q2 * q3
}
