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

type PrimitiveEquation struct {
	eq         Equation
	primitives []Term
}

type Term struct {
	termName  string
	regName   string
	regNumber int
}

func (t Term) String() string {
	return t.termName
}

func (e Equation) String() string {
	return fmt.Sprintf("%s -> %s %s %s   Primitives: ", e.dest, e.arg1, e.op, e.arg2)
	// return fmt.Sprintf("%s -> %s %s %s   Primitives: %v", e.dest, e.arg1, e.op, e.arg2, e.primitives)
}

func (p PrimitiveEquation) String() string {
	return fmt.Sprintf("%s -> %s %s %s   Primitives: %v", p.eq.dest, p.eq.arg1, p.eq.op, p.eq.arg2, p.primitives)
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

func FindIncorrectTerms(registers map[string]uint8, equations []Equation) []string {
	incorrectTerms := []string{}
	rule1 := []string{}
	rule2 := []string{}
	lastZBit := GetSpecialBits(equations)
	// firstZBit := "z00"
	for _, eq := range equations {
		if isZTerm(eq.dest) {
			if eq.op != "XOR" && eq.dest != lastZBit {
				incorrectTerms = append(incorrectTerms, eq.dest)
				rule1 = append(rule1, eq.dest)
			}
		} else {
			if eq.op == "XOR" && !(isInputBit(eq.arg1) || isInputBit(eq.arg2)) {
				incorrectTerms = append(incorrectTerms, eq.dest)
				rule2 = append(rule2, eq.dest)
			}
		}
	}

	primitiveMap := GetPrimitiveMap(equations)
	swaps := make(map[string]string)
	for _, reg := range rule2 {
		v := primitiveMap[reg]
		targetReg := getMaxPrimitiveZReg(v)
		idx, _ := lib.FindInArray(targetReg, rule1)
		if idx == -1 {
			fmt.Fprintf(os.Stderr, "couldn't find target %s in %v\n", targetReg, rule1)
		}
		swaps[targetReg] = reg
		swaps[reg] = targetReg
	}

	newEquations := []Equation{}
	for _, eq := range equations {
		target, found := swaps[eq.dest]
		newEq := eq
		if found {
			newEq.dest = target
		}
		newEquations = append(newEquations, newEq)
	}

	originalX := ParseResults(registers, "x")
	originalY := ParseResults(registers, "y")

	result := ProcessEquations(registers, newEquations)
	testZ := ParseResults(result, "z")

	actualZ := originalX + originalY
	fmt.Fprintf(os.Stderr, "origX: %d, 0%s\n", originalX, strconv.FormatInt(originalX, 2))
	fmt.Fprintf(os.Stderr, "origY: %d, 0%s\n", originalY, strconv.FormatInt(originalY, 2))
	fmt.Fprintf(os.Stderr, "actlZ: %d, %s\n", actualZ, strconv.FormatInt(actualZ, 2))
	fmt.Fprintf(os.Stderr, "testZ: %d, %s\n", testZ, strconv.FormatInt(testZ, 2))
	xorTest := actualZ ^ testZ
	fmt.Fprintf(os.Stderr, "XOR %14d, %s\n", 0, strconv.FormatInt(xorTest, 2))

	return incorrectTerms
}

func getSortedPrims(terms []Term) []string {
	prims := []string{}
	for _, t := range terms {
		prims = append(prims, t.termName)
	}
	sort.Strings(prims)
	return prims
}

func getMaxPrimitiveZReg(pEq PrimitiveEquation) string {
	maxPrim := 0
	for _, prim := range pEq.primitives {
		if prim.regNumber > maxPrim {
			maxPrim = prim.regNumber
		}
	}

	return fmt.Sprintf("z%02d", maxPrim)
}

func isInputBit(reg string) bool {
	return strings.HasPrefix(reg, "x") || strings.HasPrefix(reg, "y")
}

func GetSpecialBits(equations []Equation) string {
	lastZBit := ""

	maxZBit := -1
	for _, eq := range equations {
		if strings.HasPrefix(eq.dest, "z") {
			t := GetTerm(eq.dest)
			if t.regName == "z" && t.regNumber > maxZBit {
				lastZBit = eq.dest
				maxZBit = t.regNumber
			}
		}
	}

	return lastZBit
}

func isZTerm(reg string) bool {
	return strings.HasPrefix(reg, "z")
}

func GetPrimitiveMap(equations []Equation) map[string]PrimitiveEquation {
	// turn equation list into mapping for easy access
	eqIndex := make(map[string]PrimitiveEquation)
	for _, eq := range equations {
		eqIndex[eq.dest] = PrimitiveEquation{eq: eq, primitives: []Term{}}
	}

	// fill in all of the primitives
	for name, eq := range eqIndex {
		prims := GetPrimitives(eq.eq, &eqIndex)
		eq.primitives = prims
		eqIndex[name] = eq
	}

	return eqIndex
}

func GetPrimitives(eq Equation, eqIndex *map[string]PrimitiveEquation) []Term {
	eqPrim, found := (*eqIndex)[eq.dest]
	if found && len(eqPrim.primitives) > 0 {
		return eqPrim.primitives
	}
	primitives := []Term{}
	if !isSourceWire(eq.arg1) {
		arg1, found := (*eqIndex)[eq.arg1]
		if !found {
			fmt.Fprintf(os.Stderr, "Couldn't find %s in eqIndex\n", eq.arg1)
		}
		if len(arg1.primitives) == 0 {
			prim := GetPrimitives(arg1.eq, eqIndex)
			arg1 = (*eqIndex)[eq.arg1]
			arg1.primitives = prim
			(*eqIndex)[eq.arg1] = arg1
		}
		for _, prim := range arg1.primitives {
			primitives = lib.AppendDistinct(primitives, prim)
		}
	} else {
		primitives = append(primitives, GetTerm(eq.arg1))
	}

	if !isSourceWire(eq.arg2) {
		arg2, found := (*eqIndex)[eq.arg2]
		if !found {
			fmt.Fprintf(os.Stderr, "Couldn't find %s in eqIndex\n", eq.arg2)
		}
		if len(arg2.primitives) == 0 {
			prim := GetPrimitives(arg2.eq, eqIndex)
			arg2 = (*eqIndex)[eq.arg2]
			arg2.primitives = prim
			(*eqIndex)[eq.arg2] = arg2
		}
		for _, prim := range arg2.primitives {
			primitives = lib.AppendDistinct(primitives, prim)
		}
	} else {
		primitives = append(primitives, GetTerm(eq.arg2))
	}

	return primitives
}

func GetTerm(arg string) Term {
	regex := regexp.MustCompile("([a-zA-Z])([0-9][0-9])")
	match := regex.FindStringSubmatch(arg)
	if len(match) != 3 {
		fmt.Fprintf(os.Stderr, "Unable to parse register name: %s\n", arg)
	}
	prefix := match[1]
	regNumber, err := strconv.Atoi(match[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse register number: %s, %v\n", match[2], err)
	}
	return Term{termName: arg, regName: prefix, regNumber: regNumber}
}

func isSourceWire(arg string) bool {
	return strings.HasPrefix(arg, "x") || strings.HasPrefix(arg, "y") || strings.HasPrefix(arg, "z")
}

func PrintEquations(eq []PrimitiveEquation) {
	sb := strings.Builder{}
	sb.WriteString("[\n")

	for _, e := range eq {
		sb.WriteString("\t")
		sb.WriteString(e.String())
		sb.WriteString("\n")
	}

	sb.WriteString("]\n")

	fmt.Fprint(os.Stderr, sb.String())
}
