package main

import (
	"strings"
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimple3Part1(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")
	machine := ParseInput(content)
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "0,3,5,4,3,0")
}

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	machine := ParseInput(content)
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "4,6,3,5,6,3,5,2,1,0")
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	machine := ParseInput(content)
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "7,3,0,5,7,1,4,0,5")
}

func TestProgramSafe(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")
	machine := ParseInput(content)
	validated := machine.RunProgramSafe()
	assert.Equal(t, validated, true)
	machine = ParseInput(content)
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "0,3,5,4,3,0")
}

func TestSimpleExamplePart2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	template := ParseInput(content)
	machine := Machine{RegA: 117440, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "0,3,5,4,3,0")
}

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	template := ParseInput(content)
	newRegA := FindARegValue(template, 0)
	assert.Equal(t, newRegA, int64(117440))
	machine := Machine{RegA: newRegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "0,3,5,4,3,0")
}

func TestFastPart2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")
	template := ParseInput(content)
	newRegA := FastFindRegA(template)
	assert.Equal(t, newRegA, int64(117440))
	machine := Machine{RegA: newRegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "0,3,5,4,3,0")
}

func TestFastFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	template := ParseInput(content)
	newRegA := FastFindRegA(template)
	assert.Equal(t, newRegA, int64(202972175280682))
	machine := Machine{RegA: newRegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
	machine.RunProgram()
	output := machine.ReadOutput()
	assert.Equal(t, output, "2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0")
}

// func TestFullPart2(t *testing.T) {
// 	content := lib.ReadInput("input.txt")
// 	assert.NotEqual(t, content, "")
// 	template := ParseInput(content)
// 	newRegA := FindARegValue(template, 20000000000)
// 	assert.Equal(t, newRegA, int64(117440))
// 	machine := Machine{RegA: newRegA, RegB: template.RegB, RegC: template.RegC, Ins: 0, Cmd: template.Cmd, numOutput: 0, Out: strings.Builder{}}
// 	machine.RunProgram()
// 	output := machine.ReadOutput()
// 	assert.Equal(t, output, "0,3,5,4,3,0")
// }
