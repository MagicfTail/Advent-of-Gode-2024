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

func part1Func(input string) any {
	left, right := common.ParseDoubleListInt(input)

	common.SortIntSliceLt(left)
	common.SortIntSliceLt(right)

	diffSLice, err := common.ZipMap(right, left, common.IntDiff)
	if err != nil {
		return "mapping values failed"
	}

	total := common.Reduce(diffSLice, 0, common.IntAdd)

	return total
}

func part2Func(input string) any {
	left, right := common.ParseDoubleListInt(input)

	countMap := common.SliceToCountMap(right)

	out := 0
	for _, entry := range left {
		out += entry * countMap[entry]
	}

	return out
}
