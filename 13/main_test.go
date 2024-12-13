package main

import (
	"regexp"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	machines := ParseMachines(content)

	numPrizes, totalTokens := ProcessMachines(machines)
	assert.Equal(t, numPrizes, 2)
	assert.Equal(t, totalTokens, 480)
}

func TestSolveEquation(t *testing.T) {
	buttonA := Location{x: 94, y: 34}
	buttonB := Location{x: 22, y: 67}
	prize := Location{x: 8400, y: 5400}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.Equal(t, a, 80)
	assert.Equal(t, b, 40)
}

func TestSolveEquationNoSolution(t *testing.T) {
	buttonA := Location{x: 26, y: 66}
	buttonB := Location{x: 67, y: 21}
	prize := Location{x: 12748, y: 12176}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.Equal(t, a, -1)
	assert.Equal(t, b, -1)
}

func TestRegexDigits(t *testing.T) {
	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	loc := getRegexMatches("Button A: X+12, Y+20", re)
	assert.Equal(t, int(loc.x), 12)
	assert.Equal(t, int(loc.y), 20)
}

func TestFullPart1(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	machines := ParseMachines(content)

	numPrizes, totalTokens := ProcessMachines(machines)
	assert.Equal(t, numPrizes, 186)
	assert.Equal(t, totalTokens, 39996)
}

func TestSimplePart2(t *testing.T) {
	content := ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	machines := ParseMachinesCorrected(content)

	numPrizes, totalTokens := ProcessMachines(machines)
	assert.Equal(t, numPrizes, 2)
	assert.NotEqual(t, totalTokens, 480)
}

func TestFullPart2(t *testing.T) {
	content := ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	machines := ParseMachinesCorrected(content)

	numPrizes, totalTokens := ProcessMachines(machines)
	assert.Equal(t, numPrizes, 134)
	assert.Equal(t, totalTokens, 73267584326867)
}

func TestSolveEquationCorrectedNoSolution(t *testing.T) {
	buttonA := Location{x: 94, y: 34}
	buttonB := Location{x: 22, y: 67}
	prize := Location{x: 10000000008400, y: 10000000005400}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.Equal(t, a, -1)
	assert.Equal(t, b, -1)
}

func TestSolveEquationCorrected(t *testing.T) {
	buttonA := Location{x: 26, y: 66}
	buttonB := Location{x: 67, y: 21}
	prize := Location{x: 10000000012748, y: 10000000012176}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.NotEqual(t, a, -1)
	assert.NotEqual(t, b, -1)
}

func TestSolveEquationCorrectedNoSolution2(t *testing.T) {
	buttonA := Location{x: 17, y: 86}
	buttonB := Location{x: 84, y: 37}
	prize := Location{x: 10000000007870, y: 10000000006450}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.Equal(t, a, -1)
	assert.Equal(t, b, -1)
}

func TestSolveEquationCorrected2(t *testing.T) {
	buttonA := Location{x: 69, y: 23}
	buttonB := Location{x: 27, y: 71}
	prize := Location{x: 10000000018641, y: 10000000010279}

	a, b := SolveEquation(prize, buttonA, buttonB)
	assert.NotEqual(t, a, -1)
	assert.NotEqual(t, b, -1)
}
