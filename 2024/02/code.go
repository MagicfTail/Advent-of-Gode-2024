package main

import (
	"aoc-in-go/common"

	"github.com/jpillora/puzzler/harness/aoc"
)

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

var increasing = map[int]struct{}{1: {}, 2: {}, 3: {}}
var decreasing = map[int]struct{}{-1: {}, -2: {}, -3: {}}

func part1Func(input string) any {
	lines := common.Parse2dArrayIntSplit(input)

	total := 0

	for _, line := range lines {
		diffSlice := common.InterMap(line, common.IntSubReverse)
		diffSet := common.SliceToSet(diffSlice)

		if common.Subset(diffSet, increasing) || common.Subset(diffSet, decreasing) {
			total += 1
		}
	}

	return total
}



func part2Func(input string) any {
	lines := common.Parse2dArrayIntSplit(input)

	total := 0

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			newLine := make([]int, 0, len(line) - 1)
			for j, val := range line {
				if i == j {
					continue
				}

				newLine = append(newLine, val)
			}
			diffSlice := common.InterMap(newLine, common.IntSubReverse)
			diffSet := common.SliceToSet(diffSlice)
			
			if common.Subset(diffSet, increasing) || common.Subset(diffSet, decreasing) {
				total += 1
				break
			}
		}
	}

	return total
}
