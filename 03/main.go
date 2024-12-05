package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World")
}

const (
	StateEmpty     = iota
	StateM         = iota
	StateU         = iota
	StateL         = iota
	StateOpenParen = iota
	StateArg1      = iota
	StateComma     = iota
	StateArg2      = iota
)

type Command struct {
	cmd  string
	arg1 int
	arg2 int
}

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func ParseMult(content string) []Command {
	commands := []Command{}
	state := StateEmpty
	arg1 := strings.Builder{}
	arg2 := strings.Builder{}

	for _, c := range content {
		switch state {
		case StateEmpty:
			if c == 'm' {
				state = StateM
			}
		case StateM:
			if c == 'u' {
				state = StateU
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateU:
			if c == 'l' {
				state = StateL
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateL:
			if c == '(' {
				state = StateOpenParen
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateOpenParen:
			if isNumeric(c) {
				state = StateArg1
				arg1.WriteRune(c)
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateArg1:
			if isNumeric(c) {
				arg1.WriteRune(c)
				continue
			}
			if c == ',' {
				state = StateComma
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateComma:
			if isNumeric(c) {
				arg2.WriteRune(c)
				state = StateArg2
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		case StateArg2:
			if isNumeric(c) {
				arg2.WriteRune(c)
				continue
			}
			if c == ')' {
				a1, err := strconv.Atoi(arg1.String())
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse arg1 value %s: %v\n", arg1.String(), err)
				}

				a2, err := strconv.Atoi(arg2.String())
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse arg2 value %s: %v\n", arg2.String(), err)
				}
				cmd := Command{cmd: "mul", arg1: a1, arg2: a2}
				arg1.Reset()
				arg2.Reset()
				state = StateEmpty
				commands = append(commands, cmd)
				fmt.Fprintf(os.Stderr, "Found %d, %d\n", a1, a2)
			} else {
				state = StateEmpty
				arg1.Reset()
				arg2.Reset()
			}
		}
	}
	return commands
}

func ProcessCommands(commands []Command) int {
	sum := 0
	for _, cmd := range commands {
		sum += cmd.arg1 * cmd.arg2
	}

	return sum
}

func isNumeric(c rune) bool {
	return c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' || c == '0'
}
