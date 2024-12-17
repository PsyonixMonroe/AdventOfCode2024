package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Machine struct {
	RegA      int64
	RegB      int64
	RegC      int64
	Ins       int
	Cmd       []uint8
	numOutput int
	Out       strings.Builder
}

func (m Machine) ReadOutput() string {
	return m.Out.String()
}

func (m *Machine) incIns() {
	m.Ins += 2
}

func (m Machine) getOpCode() uint8 {
	return m.Cmd[m.Ins]
}

func (m Machine) getOperand() uint8 {
	if m.Ins >= len(m.Cmd)-1 {
		fmt.Fprintf(os.Stderr, "Trying to read operand past cmds\n")
	}
	return m.Cmd[m.Ins+1]
}

func (m Machine) getComboOperandValue(operand uint8) int64 {
	if operand <= 3 {
		return int64(operand)
	}
	if operand == 4 {
		return m.RegA
	}
	if operand == 5 {
		return m.RegB
	}
	if operand == 6 {
		return m.RegC
	}
	fmt.Fprintf(os.Stderr, "Unknown Combo Operand: %d\n", operand)
	return -1
}

// Perform the dv operation and return the result
func (m Machine) dv(operand uint8) int64 {
	cmb := m.getComboOperandValue(operand)
	num := m.RegA
	den := int64(1)
	for range cmb {
		den *= 2
	}

	return num / den
}

// adv
// The adv instruction (opcode 0) performs division.
// The numerator is the value in the A register.
// The denominator is found by raising 2 to the power of the instruction's combo operand.
// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
// The result of the division operation is truncated to an integer and then written to the A register.
func (m *Machine) adv(operand uint8) {
	m.RegA = m.dv(operand)
	m.incIns()
}

// bdv
// The bdv instruction (opcode 6) works exactly like the adv instruction except that
// the result is stored in the B register. (The numerator is still read from the A register.)
func (m *Machine) bdv(operand uint8) {
	m.RegB = m.dv(operand)
	m.incIns()
}

// cdv
// The cdv instruction (opcode 7) works exactly like the adv instruction except that
// the result is stored in the C register. (The numerator is still read from the A register.)
func (m *Machine) cdv(operand uint8) {
	m.RegC = m.dv(operand)
	m.incIns()
}

// bxl
// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and
// the instruction's literal operand, then stores the result in register B.
func (m *Machine) bxl(operand uint8) {
	m.RegB = m.RegB ^ int64(operand)
	m.incIns()
}

// bst
// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8
// (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func (m *Machine) bst(operand uint8) {
	cmb := m.getComboOperandValue(operand)
	m.RegB = cmb % 8
	m.incIns()
}

// jnz
// The jnz instruction (opcode 3) does nothing if the A register is 0.
// However, if the A register is not zero, it jumps by setting the instruction pointer
// to the value of its literal operand;
// if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func (m *Machine) jnz(operand uint8) {
	if m.RegA == 0 {
		m.incIns()
		return
	}
	m.Ins = int(operand)
}

// bxc
// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
// then stores the result in register B.
// (For legacy reasons, this instruction reads an operand but ignores it.)
func (m *Machine) bxc(operand uint8) {
	m.RegB = m.RegB ^ m.RegC
	_ = operand
	m.incIns()
}

// out
// The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
// then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func (m *Machine) out(operand uint8) {
	cmb := m.getComboOperandValue(operand)
	writeVal := cmb % 8
	if m.Out.Len() > 0 {
		m.Out.WriteString(",")
	}
	m.Out.WriteString(fmt.Sprintf("%d", writeVal))
	m.incIns()
}

func (m *Machine) outSolver(operand uint8) bool {
	if m.numOutput >= len(m.Cmd) {
		fmt.Fprintf(os.Stderr, "Output past end of Cmd\n")
		return false
	}
	cmb := m.getComboOperandValue(operand)
	outVal := cmb % 8
	result := m.Cmd[m.numOutput] == uint8(outVal)
	m.numOutput++
	m.out(operand)
	return result
}

func (m *Machine) RunProgram() {
	for m.Ins < len(m.Cmd) {
		opCode := m.getOpCode()
		operand := m.getOperand()
		switch opCode {
		case 0:
			m.adv(operand)
		case 1:
			m.bxl(operand)
		case 2:
			m.bst(operand)
		case 3:
			m.jnz(operand)
		case 4:
			m.bxc(operand)
		case 5:
			m.out(operand)
		case 6:
			m.bdv(operand)
		case 7:
			m.cdv(operand)
		default:
			fmt.Fprintf(os.Stderr, "Unknown opCode: %d\n", opCode)
		}
	}
}

func (m Machine) CmdAsOutput() string {
	out := strings.Builder{}
	for i := range len(m.Cmd) {
		if out.Len() > 0 {
			out.WriteString(",")
		}
		out.WriteString(fmt.Sprintf("%d", m.Cmd[i]))
	}

	return out.String()
}

// FindARegValue brute force solution, probably will never finish
func FindARegValue(template Machine, startval int64) int64 {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(12)
	solution := make(chan int64)
	expectedOutput := template.CmdAsOutput()
	for i := range 12 {
		go func(num int64) {
			defer wg.Done()
			regA := startval + num
			for {
				if (regA-num)%50000000 == 0 && regA != num {
					fmt.Fprintf(os.Stderr, "[%d] Processing %d\n", num, regA)
				}
				select {
				case <-ctx.Done():
					return
				default:
					success := testRegAVal(regA, template, expectedOutput)
					if success {
						solution <- regA
						return
					}
				}
				regA += 12
			}
		}(int64(i + 1))
	}

	result := <-solution

	minVal := int64(math.MaxInt64)
	for i := range 20 {
		regA := result - int64(i)
		success := testRegAVal(regA, template, expectedOutput)
		if success && regA < minVal {
			minVal = regA
		}
	}

	cancel()
	wg.Wait()

	return minVal
}

type Job struct {
	regA     int64
	position int
}

func FastFindRegA(template Machine) int64 {
	expectedOutput := template.CmdAsOutput()
	work := lib.NewQueue[Job]()
	work.Enqueue(Job{regA: 0, position: 0})
	for !work.IsEmpty() {
		job := work.Dequeue().Get()
		regA := job.regA * 0o10
		position := job.position
		newExpected := expectedOutput[len(expectedOutput)-(1+position*2):]
		for i := range int64(8) {
			success := testRegAValNoSuccess(regA+i, template, newExpected)
			if success {
				finalTest := testRegAVal(regA+i, template, expectedOutput)
				if finalTest {
					return regA + i
				} else {
					work.Enqueue(Job{regA: regA + i, position: position + 1})
				}
			}
		}
	}

	return -1
}

func testRegAValNoSuccess(RegA int64, template Machine, expectedOutput string) bool {
	machine := Machine{RegA: RegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	machine.RunProgram()
	output := machine.ReadOutput()

	return output == expectedOutput
}

func testRegAVal(RegA int64, template Machine, expectedOutput string) bool {
	machine := Machine{RegA: RegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	success := machine.RunProgramSafe()
	output := machine.ReadOutput()

	return success && output == expectedOutput
}

func (m *Machine) RunProgramSafe() bool {
	for m.Ins < len(m.Cmd) {
		opCode := m.getOpCode()
		operand := m.getOperand()
		switch opCode {
		case 0:
			m.adv(operand)
		case 1:
			m.bxl(operand)
		case 2:
			m.bst(operand)
		case 3:
			m.jnz(operand)
		case 4:
			m.bxc(operand)
		case 5:
			correct := m.outSolver(operand)
			if !correct {
				return false
			}
		case 6:
			m.bdv(operand)
		case 7:
			m.cdv(operand)
		default:
			fmt.Fprintf(os.Stderr, "Unknown opCode: %d\n", opCode)
		}
	}

	return true
}

func ParseInput(content string) Machine {
	lines := strings.Split(strings.Trim(content, "\n"), "\n")
	if len(lines) != 5 {
		fmt.Fprintf(os.Stderr, "Unable to parse machines from content\n===\n%s\n===\n", content)
		return Machine{}
	}

	regA := GetMatchInt(`Register A: (\d+)`, lines[0])
	regB := GetMatchInt(`Register B: (\d+)`, lines[1])
	regC := GetMatchInt(`Register C: (\d+)`, lines[2])

	programString := GetString(`Program: (.*)$`, lines[4])

	program := []uint8{}
	for _, val := range strings.Split(strings.Trim(programString, "\n\t "), ",") {
		input, err := strconv.Atoi(val)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unknown Value: %s\n", val)
			continue
		}

		program = append(program, uint8(input))
	}

	return Machine{RegA: regA, RegB: regB, RegC: regC, Ins: 0, Cmd: program, Out: strings.Builder{}}
}

func GetString(reg string, line string) string {
	regex := regexp.MustCompile(reg)
	match := regex.FindStringSubmatch(line)
	if len(match) != 2 {
		fmt.Fprintf(os.Stderr, "Unable to find match with %s in %s\n", reg, line)
		return "-1"
	}

	return match[1]
}

func GetMatchInt(reg string, line string) int64 {
	stringVal := GetString(reg, line)
	val, err := strconv.Atoi(stringVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse match with %s\n%v", stringVal, err)
		return -1
	}

	return int64(val)
}
