package main

import (
	"aoc-in-go/common"
	"os"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	b, _ := os.ReadFile("input-example.txt")
	part1Func(string(b))
	aoc.Harness(run)
}

var paths = map[fromTo]string{
	{'0', '0'}: "A",
	{'1', '1'}: "A",
	{'2', '2'}: "A",
	{'3', '3'}: "A",
	{'4', '4'}: "A",
	{'5', '5'}: "A",
	{'6', '6'}: "A",
	{'7', '7'}: "A",
	{'8', '8'}: "A",
	{'9', '9'}: "A",
	{'A', 'A'}: "A",
	{'^', '^'}: "A",
	{'<', '<'}: "A",
	{'>', '>'}: "A",
	{'v', 'v'}: "A",
	{'A', '0'}: "<A",
	{'0', 'A'}: ">A",
	{'A', '1'}: "^<<A",
	{'1', 'A'}: ">>vA",
	{'A', '2'}: "<^A",
	{'2', 'A'}: "v>A",
	{'A', '3'}: "^A",
	{'3', 'A'}: "vA",
	{'A', '4'}: "^^<<A",
	{'4', 'A'}: ">>vvA",
	{'A', '5'}: "<^^A",
	{'5', 'A'}: "vv>A",
	{'A', '6'}: "^^A",
	{'6', 'A'}: "vvA",
	{'A', '7'}: "^^^<<A",
	{'7', 'A'}: ">>vvvA",
	{'A', '8'}: "<^^^A",
	{'8', 'A'}: "vvv>A",
	{'A', '9'}: "^^^A",
	{'9', 'A'}: "vvvA",
	{'0', '1'}: "^<A",
	{'1', '0'}: ">vA",
	{'0', '2'}: "^A",
	{'2', '0'}: "vA",
	{'0', '3'}: "^>A",
	{'3', '0'}: "<vA",
	{'0', '4'}: "^<^A",
	{'4', '0'}: ">vvA",
	{'0', '5'}: "^^A",
	{'5', '0'}: "vvA",
	{'0', '6'}: "^^>A",
	{'6', '0'}: "<vvA",
	{'0', '7'}: "^^^<A",
	{'7', '0'}: ">vvvA",
	{'0', '8'}: "^^^A",
	{'8', '0'}: "vvvA",
	{'0', '9'}: "^^^>A",
	{'9', '0'}: "<vvvA",
	{'1', '2'}: ">A",
	{'2', '1'}: "<A",
	{'1', '3'}: ">>A",
	{'3', '1'}: "<<A",
	{'1', '4'}: "^A",
	{'4', '1'}: "vA",
	{'1', '5'}: "^>A",
	{'5', '1'}: "<vA",
	{'1', '6'}: "^>>A",
	{'6', '1'}: "<<vA",
	{'1', '7'}: "^^A",
	{'7', '1'}: "vvA",
	{'1', '8'}: "^^>A",
	{'8', '1'}: "<vvA",
	{'1', '9'}: "^^>>A",
	{'9', '1'}: "<<vvA",
	{'2', '3'}: ">A",
	{'3', '2'}: "<A",
	{'2', '4'}: "<^A",
	{'4', '2'}: "v>A",
	{'2', '5'}: "^A",
	{'5', '2'}: "vA",
	{'2', '6'}: "^>A",
	{'6', '2'}: "<vA",
	{'2', '7'}: "<^^A",
	{'7', '2'}: "vv>A",
	{'2', '8'}: "^^A",
	{'8', '2'}: "vvA",
	{'2', '9'}: "^^>A",
	{'9', '2'}: "<vvA",
	{'3', '4'}: "<<^A",
	{'4', '3'}: "v>>A",
	{'3', '5'}: "<^A",
	{'5', '3'}: "v>A",
	{'3', '6'}: "^A",
	{'6', '3'}: "vA",
	{'3', '7'}: "<<^^A",
	{'7', '3'}: "vv>>A",
	{'3', '8'}: "<^^A",
	{'8', '3'}: "vv>A",
	{'3', '9'}: "^^A",
	{'9', '3'}: "vvA",
	{'4', '5'}: ">A",
	{'5', '4'}: "<A",
	{'4', '6'}: ">>A",
	{'6', '4'}: "<<A",
	{'4', '7'}: "^A",
	{'7', '4'}: "vA",
	{'4', '8'}: "^>A",
	{'8', '4'}: "<vA",
	{'4', '9'}: "^>>A",
	{'9', '4'}: "<<vA",
	{'5', '6'}: ">A",
	{'6', '5'}: "<A",
	{'5', '7'}: "<^A",
	{'7', '5'}: "v>A",
	{'5', '8'}: "^A",
	{'8', '5'}: "vA",
	{'5', '9'}: "^>A",
	{'9', '5'}: "<vA",
	{'6', '7'}: "<<^A",
	{'7', '6'}: "v>>A",
	{'6', '8'}: "<^A",
	{'8', '6'}: "v>A",
	{'6', '9'}: "^A",
	{'9', '6'}: "vA",
	{'7', '8'}: ">A",
	{'8', '7'}: "<A",
	{'7', '9'}: ">>A",
	{'9', '7'}: "<<A",
	{'8', '9'}: ">A",
	{'9', '8'}: "<A",
	{'<', '^'}: ">^A",
	{'^', '<'}: "v<A",
	{'<', 'v'}: ">A",
	{'v', '<'}: "<A",
	{'<', '>'}: ">>A",
	{'>', '<'}: "<<A",
	{'<', 'A'}: ">>^A",
	{'A', '<'}: "v<<A",
	{'^', 'v'}: "vA",
	{'v', '^'}: "^A",
	{'^', '>'}: "v>A",
	{'>', '^'}: "<^A",
	{'^', 'A'}: ">A",
	{'A', '^'}: "<A",
	{'v', '>'}: ">A",
	{'>', 'v'}: "<A",
	{'v', 'A'}: "^>A",
	{'A', 'v'}: "<vA",
	{'>', 'A'}: "^A",
	{'A', '>'}: "vA",
}

type fromTo struct {
	from rune
	to rune
}

type fromToDepth struct {
	fromTo fromTo
	depth int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2Func(input)
	}

	return part1Func(input)
}


func part1Func(input string) any {
	lines := strings.Split(input, "\n")

	seen := make(map[fromToDepth]int)

	total := 0
	for _, pad1Code := range lines {
		numVal := common.AtoI(pad1Code[:3])
		start := 'A'
		shortest := 0
		for i := range len(pad1Code) {
			next := rune(pad1Code[i])
			buttons := fromTo{start, next}
			shortest += handleFromTo(seen, buttons, 2)
			start = next
		}
		total += shortest * numVal
	}

	return total
}


func handleFromTo(seen map[fromToDepth]int, current fromTo, depth int) int {
	if depth == 0 {
		return len(paths[current])
	}
	currentInDepth := fromToDepth{fromTo: current, depth: depth}

	if val, ok := seen[currentInDepth]; ok {
		return val
	}

	needed := paths[current]
	buttons := 0

	start := 'A'
	for i := range len(needed) {
		next := rune(needed[i])

		currentNeeded := fromTo{start, next}
		buttons += handleFromTo(seen, currentNeeded, depth - 1)
		start = next
	}
	seen[currentInDepth] = buttons

	return buttons
}

func part2Func(input string) any {
	lines := strings.Split(input, "\n")

	seen := make(map[fromToDepth]int)

	total := 0
	for _, pad1Code := range lines {
		numVal := common.AtoI(pad1Code[:3])
		start := 'A'
		shortest := 0
		for i := range len(pad1Code) {
			next := rune(pad1Code[i])
			buttons := fromTo{start, next}
			shortest += handleFromTo(seen, buttons, 25)
			start = next
		}
		total += shortest * numVal
	}

	return total
}
