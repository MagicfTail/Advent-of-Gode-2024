package main

import (
	"aoc-in-go/common"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const inputRegex = `Register A: (\d*)
Register B: (\d*)
Register C: (\d*)

Program: (.*)`

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	if part2 {
		return part2Func(input)
	}

	return part1Func(input)
}

func parseInput(input string) (a, b, c int, program []int) {
	re := regexp.MustCompile(inputRegex)
	found := re.FindStringSubmatch(input)

	a = common.AtoI(found[1])
	b = common.AtoI(found[2])
	c = common.AtoI(found[3])

	stringProgram := strings.Split(found[4], ",")
	program = common.MapIgnoreErr(stringProgram, strconv.Atoi)

	return
}

func getComboValue(a, b, c, op int) int {
	switch op {
	case 7:
		panic("Invalid op")
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return op
	}
}

func part1Func(input string) any {
	a, b, c, program := parseInput(input)

	out := []int{}

	pc := 0

	for pc < len(program) {
		ins  := program[pc]
		op := program[pc + 1]

		switch ins {
		case 0:
			a = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		case 1:
			b = b ^ op
		case 2:
			b = getComboValue(a, b, c, op) % 8
		case 3:
			if a != 0 {
				pc = op
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			out = append(out, getComboValue(a, b, c, op) % 8)
		case 6:
			b = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		case 7:
			c = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		}

		pc+= 2
	}
	
	strOut := common.Map(out, strconv.Itoa)

	return strings.Join(strOut, ",")
}

// Hardcoded for this specific
func part2Func(input string) any {
	_, _, _, program := parseInput(input)
	target := strings.Join(common.Map(program, strconv.Itoa), "")

	_, aVal := depthFirst2(target, 0, 1, program)

	return aVal
}

func computeOut2(a int, program []int) string {
	str := ""

	b := 0
	c := 0

	pc := 0

	for pc < len(program) {
		ins  := program[pc]
		op := program[pc + 1]

		switch ins {
		case 0:
			a = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		case 1:
			b = b ^ op
		case 2:
			b = getComboValue(a, b, c, op) % 8
		case 3:
			if a != 0 {
				pc = op
				continue
			}
		case 4:
			b = b ^ c
		case 5:
			str+= strconv.Itoa(getComboValue(a, b, c, op) % 8)
		case 6:
			b = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		case 7:
			c = a / int(math.Pow(2, float64(getComboValue(a, b, c, op))))
		}

		pc+= 2
	}

	return str
}

func depthFirst2(target string, depth, current int, program []int) (bool, int) {
	for i := range 8 {
		new := current + i

		str := computeOut2(new, program)

		ok, done := str == target[len(target) - len(str):], len(target) == len(str)
		if done && ok {
			return true, new
		}

		newCurrent := new * 8

		if ok {
			if done, val := depthFirst2(target, depth + 1, newCurrent, program); done {
				return done, val
			}
		}
	}

	return false, -1
}

// Faster but non-generic solution
func computeOut(a int) string {
	str := ""

	b := 0
	c := 0

	for a > 0 {
		b = a % 8
		b = b ^ 2
		c = a / (1 << b)
		b = b ^ 3
		b = b ^ c 
		a = a >> 3
		str += strconv.Itoa(b % 8)
	}

	return str
}

func depthFirst(target string, depth, current int) (bool, int) {
	for i := range 8 {
		new := current + i

		str := computeOut(new)

		ok, done := str == target[len(target) - len(str):], len(target) == len(str)
		if done && ok {
			return true, new
		}

		newCurrent := new * 8

		if ok {
			if done, val := depthFirst(target, depth + 1, newCurrent); done {
				return done, val
			}
		}
	}

	return false, -1
}
