package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(2024))
}

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(4))
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	registers, equations := ParseInput(content)
	registers = ProcessEquations(registers, equations)
	finalBuffer := ParseResults(registers, "z")

	assert.Equal(t, finalBuffer, int64(43559017878162))
}

// func TestSimplePart2(t *testing.T) {
// 	content := lib.ReadInput("test3.txt")
// 	assert.NotEqual(t, content, "")
//
// 	_, equations := ParseInput(content)
// 	result := FindIncorrectTerms(equations)
//
// 	_ = result
// }

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	registers, equations := ParseInput(content)
	result := FindIncorrectTerms(registers, equations)

	fmt.Fprintf(os.Stderr, "%v\n", result)
}

func TestFullCorrectedPart2(t *testing.T) {
	content := lib.ReadInput("input_fixed.txt")
	assert.NotEqual(t, content, "")

	registers, equations := ParseInput(content)
	result := FindIncorrectTerms(registers, equations)
	result = append(result, "mwh")
	result = append(result, "ggt")

	sort.Strings(result)

	final := strings.Join(result, ",")

	fmt.Fprintf(os.Stderr, "%s\n", final)
}
