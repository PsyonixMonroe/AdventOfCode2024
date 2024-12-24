package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Equation struct {
	arg1 string
	arg2 string
	op   string
	dest string
}

func ParseInput(content string) (map[string]uint8, []Equation) {
	registers := make(map[string]uint8)
	equations := []Equation{}
	lines := strings.Split(strings.Trim(content, "\n \t"), "\n")
	valRegex := regexp.MustCompile(`^([a-zA-Z0-9]+): ([\d]+)$`)
	eqRegex := regexp.MustCompile(`^([a-zA-Z0-9]+) ([a-zA-Z0-9]+) ([a-zA-Z0-9]+) -> ([a-zA-Z0-9]+)$`)
	for _, s := range lines {
		line := strings.Trim(s, "\n \t")
		if line == "" {
			continue
		}
		if strings.Index(line, ":") > 0 {
			// process value
			match := valRegex.FindStringSubmatch(line)
			if len(match) != 3 {
				fmt.Fprintf(os.Stderr, "Unable to parse value line: %s\n", line)
				continue
			}
			parsed, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't parse value from match: %s\n", match[2])
				continue
			}

			registers[match[1]] = uint8(parsed)
		} else {
			// process Equation
			match := eqRegex.FindStringSubmatch(line)
			if len(match) != 5 {
				fmt.Fprintf(os.Stderr, "Unable to parse equation line: %s\n", line)
				continue
			}
			eq := Equation{arg1: match[1], arg2: match[3], op: match[2], dest: match[4]}
			equations = append(equations, eq)
		}
	}

	return registers, equations
}

func ComputeEquation(eq Equation, registers map[string]uint8) (uint8, error) {
	a1, found := registers[eq.arg1]
	if !found {
		return 2, fmt.Errorf("could not find arg1 (%s) for equation", eq.arg1)
	}
	a2, found := registers[eq.arg2]
	if !found {
		return 2, fmt.Errorf("could not find arg2 (%s) for equation", eq.arg2)
	}

	switch eq.op {
	case "AND":
		return and(a1, a2), nil
	case "OR":
		return or(a1, a2), nil
	case "XOR":
		return xor(a1, a2), nil
	default:
		return 2, fmt.Errorf("unknown operations: %s", eq.op)
	}
}

func and(a1 uint8, a2 uint8) uint8 {
	if a1 == 1 && a2 == 1 {
		return 1
	}
	return 0
}

func or(a1 uint8, a2 uint8) uint8 {
	if a1 == 1 || a2 == 1 {
		return 1
	}
	return 0
}

func xor(a1 uint8, a2 uint8) uint8 {
	if (a1 == 0 && a2 == 1) || (a1 == 1 && a2 == 0) {
		return 1
	}
	return 0
}

func ProcessEquations(registers map[string]uint8, equations []Equation) map[string]uint8 {
	processed := 0
	for len(equations) > 0 || processed > 0 {
		toRemove := []Equation{}
		processed = 0
		for _, eq := range equations {
			if isInMap(eq.arg1, registers) && isInMap(eq.arg2, registers) {
				result, err := ComputeEquation(eq, registers)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to compute equation %v, %v\n", eq, err)
					continue
				}
				registers[eq.dest] = result
				toRemove = append(toRemove, eq)
				processed++
			}
		}
		for _, eq := range toRemove {
			equations, _ = lib.RemoveItem(equations, eq)
		}
	}

	if len(equations) > 0 {
		fmt.Fprintf(os.Stderr, "Could not complete %d equations\n", len(equations))
	}
	return registers
}

func isInMap(k string, m map[string]uint8) bool {
	_, found := m[k]
	return found
}

func ParseResults(registers map[string]uint8, bufferPrefix string) int64 {
	resultRegisters := []string{}
	for k := range registers {
		if strings.HasPrefix(k, bufferPrefix) {
			resultRegisters = append(resultRegisters, k)
		}
	}

	sort.Strings(resultRegisters)
	finalBuffer := ""
	for _, regName := range resultRegisters {
		val := registers[regName]
		finalBuffer = fmt.Sprintf("%d%s", val, finalBuffer)
	}

	result, err := strconv.ParseInt(finalBuffer, 2, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse result buffer: %s, %v\n", finalBuffer, err)
		return -1
	}
	return result
}

func ReverseArray(a []string) []string {
	s := []string{}
	for i := range a {
		s = append(s, a[len(a)-1-i])
	}

	return s
}
