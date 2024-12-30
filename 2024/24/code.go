package main

import (
	"aoc-in-go/common"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var initialRegexExp = `(.*): (\d)`
var wireRegexExp = `(.*) (.*) (.*) -> (.*)`

func main() {
	b, _ := os.ReadFile("input-example.txt")
	part2Func(string(b))
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

type gate struct {
	left string
	right string
	op string
}

func parseInput(input string) (map[string]bool, map[string]gate) {
	known := make(map[string]bool)
	gates := make(map[string]gate)

	initialRegex := regexp.MustCompile(initialRegexExp)
	wireRegex := regexp.MustCompile(wireRegexExp)
	
	initials := initialRegex.FindAllStringSubmatch(input, -1)
	wires := wireRegex.FindAllStringSubmatch(input, -1)

	for _, intital := range initials {
		wire := intital[1]
		val := intital[2]

		on := val == "1"

		known[wire] = on
	}

	for _, wire := range wires {
		in1 := wire[1]
		op := wire[2]
		in2 := wire[3]
		out := wire[4]

		gates[out] = gate{left: in1, right: in2, op: op}
	}

	return known, gates
}

func handleOp(left, right bool, op string) bool {
	switch op {
	case "XOR":
		return left != right
	case "OR":
		return left || right
	case "AND":
		return left && right
	default:
		panic(op)
	}
}

func part1Func(input string) any {
	known, gates := parseInput(input)

	total := 0
	for wire := range gates {
		if strings.HasPrefix(wire, "z") {
			number := common.AtoI(wire[1:])
			if findVal(known, gates, wire) {
				total += (1 << number)
			}
		}
	}

	return total
}

func findVal(known map[string]bool, gates map[string]gate, wire string) bool {
	if val, ok := known[wire]; ok {
		return val
	}

	gateStruct := gates[wire]

	leftVal := findVal(known, gates, gateStruct.left)
	rightVal := findVal(known, gates, gateStruct.right)

	val := handleOp(leftVal, rightVal, gateStruct.op)
	known[wire] = val
	
	return val
}

func notDefault(wire string) bool {
	return wire[0] != 'x' && wire[0] != 'y' && wire[0] != 'z'
}

func part2Func(input string) any {
	_, gates := parseInput(input)

	wrongs := common.Set[string]{}

	highestZ := "z00"

	for wire := range gates {
		if wire[0] != 'z' {
			continue
		}

		if wire > highestZ {
			highestZ = wire
		}
	}

	for wire, gate := range gates {
		if wire[0] == 'z' && gate.op != "XOR" && wire != highestZ {
			wrongs[wire] = struct{}{}
			continue
		}

		if gate.op == "XOR" && notDefault(wire) && notDefault(gate.left)  && notDefault(gate.right) {
			wrongs[wire] = struct{}{}
			continue
		}

		if gate.op == "AND" && gate.left != "x00" && gate.right != "x00" {
			for _, subGate := range gates {
				if (wire == subGate.left || wire == subGate.right) && subGate.op != "OR" {
					wrongs[wire] = struct{}{}
				}
			}
			continue
		}
		
		if gate.op == "XOR" {
			for _, subGate := range gates {
				if (wire == subGate.left || wire == subGate.right) && subGate.op == "OR" {
					wrongs[wire] = struct{}{}
				}
			}
			continue
		}
	}

	wrongList := common.SetToSlice(wrongs)

	sort.Slice(wrongList, func(i, j int) bool {
		return wrongList[i] < wrongList[j]
	})

	return strings.Join(wrongList, ",")
}
