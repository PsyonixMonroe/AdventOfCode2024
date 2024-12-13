package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

type Location struct {
	x float64
	y float64
}

type Machine struct {
	buttonA Location
	buttonB Location
	prize   Location
}

func ParseMachines(content string) []Machine {
	machines := []Machine{}
	inLines := strings.Split(content, "\n")
	buttonARegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	buttonBRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	for len(inLines) > 0 {
		if strings.TrimRightFunc(inLines[0], unicode.IsSpace) == "" {
			inLines = inLines[1:]
			continue
		}
		if len(inLines) < 3 {
			break
		}
		buttonALine := strings.Trim(inLines[0], "\n")
		buttonBLine := strings.Trim(inLines[1], "\n")
		prizeLine := strings.Trim(inLines[2], "\n")
		buttonA := getRegexMatches(buttonALine, buttonARegex)
		buttonB := getRegexMatches(buttonBLine, buttonBRegex)
		prize := getRegexMatches(prizeLine, prizeRegex)
		machines = append(machines, Machine{buttonA: buttonA, buttonB: buttonB, prize: prize})
		inLines = inLines[3:]
	}

	return machines
}

func getRegexMatches(line string, regex *regexp.Regexp) Location {
	match := regex.FindStringSubmatch(line)
	if len(match) != 3 {
		fmt.Fprintf(os.Stderr, "Unknown format for line: %s\n", line)
		return Location{x: -1, y: -1}
	}
	x := lib.StringToInt(match[1])
	y := lib.StringToInt(match[2])
	return Location{x: float64(x), y: float64(y)}
}

func ProcessMachines(machines []Machine) (int, int) {
	numPrizes := 0
	numTokens := 0
	for _, machine := range machines {
		a, b := SolveEquation(machine.prize, machine.buttonA, machine.buttonB)
		if a == -1 || b == -1 {
			continue
		}
		numPrizes++
		numTokens += a*3 + b
	}

	return numPrizes, numTokens
}

// SolveEquation function
//
// take in the two buttons and the prize location, compute the number of button presses
// to get the cursor to location. returned as (a, b)
func SolveEquation(prize Location, buttonA Location, buttonB Location) (int, int) {
	top := prize.x - buttonA.x*prize.y/buttonA.y
	bottom := buttonB.x - (buttonA.x * buttonB.y / buttonA.y)

	b := int(math.Round(top / bottom))
	a := int(math.Round((prize.y - float64(b)*buttonB.y) / buttonA.y))

	checkX := float64(a)*buttonA.x + float64(b)*buttonB.x
	checkY := float64(a)*buttonA.y + float64(b)*buttonB.y

	if checkX != prize.x || checkY != prize.y || a > 100 || b > 100 {
		return -1, -1
	}

	return a, b
}
