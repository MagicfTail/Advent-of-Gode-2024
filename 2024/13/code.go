package main

import (
	"aoc-in-go/common"
	"math"
	"regexp"

	"github.com/jpillora/puzzler/harness/aoc"
)

const inputRegex = `Button A: X\+(\d*), Y\+(\d*)
Button B: X\+(\d*), Y\+(\d*)
Prize: X=(\d*), Y=(\d*)`

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

func parseInput(input string) []entry {
	re := regexp.MustCompile(inputRegex)
	found := re.FindAllStringSubmatch(input, -1)

	entries := make([]entry, 0, len(found))

	for _, example := range found {
		entries = append(entries, entry{
			ButtonA: common.Coordinate{
				X: common.AtoI(example[1]),
				Y: common.AtoI(example[2]),
			},
			ButtonB: common.Coordinate{
				X: common.AtoI(example[3]),
				Y: common.AtoI(example[4]),
			},
			Target:  common.Coordinate{
				X: common.AtoI(example[5]),
				Y: common.AtoI(example[6]),
			},
		})
	}

	return entries
}

type entry struct {
	ButtonA common.Coordinate
	ButtonB common.Coordinate
	Target common.Coordinate
}

func part1Func(input string) any {
	entries := parseInput(input)

	total := 0

	for _, entry := range entries {
		minFound := math.MaxInt

		for a := range 100 {
			for b := range 100 {
				if a * entry.ButtonA.X + b * entry.ButtonB.X == entry.Target.X && a * entry.ButtonA.Y + b * entry.ButtonB.Y == entry.Target.Y {
					minFound = common.IntMin(minFound, 3*a + b)
				}
			}
		}

		if minFound != math.MaxInt {
			total += minFound
		}
	}

	return total
}

func part2Func(input string) any {
	entries := parseInput(input)

	offset := 10000000000000

	total := 0

	for _, entry := range entries {
		entry.Target.X += offset
		entry.Target.Y += offset
		
		x1 := (entry.ButtonB.Y * entry.Target.X - entry.ButtonB.X * entry.Target.Y) / (entry.ButtonB.Y * entry.ButtonA.X - entry.ButtonB.X * entry.ButtonA.Y)
		x2 := (entry.Target.X - entry.ButtonA.X * x1)/entry.ButtonB.X

		if x1 * entry.ButtonA.X + x2 *entry.ButtonB.X == entry.Target.X && x1 * entry.ButtonA.Y + x2 * entry.ButtonB.Y == entry.Target.Y {
			total += x1 * 3 + x2
		}
	}

	return total
}
