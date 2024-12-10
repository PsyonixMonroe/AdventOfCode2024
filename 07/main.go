package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	solution int
	args     []int
}

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func ParseInput(content string) []Equation {
	equations := []Equation{}
	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		eq := strings.Trim(line, "\n")
		splits := strings.Split(eq, ":")
		if len(splits) != 2 {
			fmt.Fprintf(os.Stderr, "Incorrect format for line: %s\n", eq)
			continue
		}
		solution, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse solution: %v\n", err)
			continue
		}

		args := []int{}
		for _, val := range strings.Fields(splits[1]) {
			arg, err := strconv.Atoi(val)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse argument: %v\n", err)
				continue
			}
			args = append(args, arg)
		}

		equations = append(equations, Equation{solution: solution, args: args})
	}

	return equations
}

func SumGoodEquations(equations []Equation) int {
	sum := 0
	for _, equation := range equations {
		sum += ProcessEquation(equation.solution, equation.args[0], equation.args[1:])
	}

	return sum
}

func ProcessEquation(solution int, lhs int, args []int) int {
	if len(args) == 0 {
		if lhs == solution {
			return solution
		} else {
			return 0
		}
	}

	var mulSol int
	var addSol int
	if len(args) == 1 {
		mulSol = lhs * args[0]
		addSol = lhs + args[0]
	} else {
		mulSol = ProcessEquation(solution, lhs*args[0], args[1:])
		addSol = ProcessEquation(solution, lhs+args[0], args[1:])
	}

	if mulSol == solution {
		return mulSol
	}

	if addSol == solution {
		return addSol
	}

	return 0
}
