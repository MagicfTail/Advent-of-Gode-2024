package main

import (
	"aoc-in-go/common"
	"strconv"
	"strings"

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
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		cases := strings.Split(line, ":")
		left := common.AtoI(cases[0])
		right := strings.Split(strings.TrimSpace(cases[1]), " ")

		numbers := make([]int, 0, len(right))

		for _, entry := range right {
			number := common.AtoI(entry)
			numbers = append(numbers, number)
		}

		if brute(left, numbers[0], numbers[1:]) {
			total += left
		}
	}

	return total
}

func brute(target, total int, entriesLeft []int) bool {
	if len(entriesLeft) == 0 {
		return target == total
	}

	mul := total * entriesLeft[0]
	add := total + entriesLeft[0]

	return brute(target, mul, entriesLeft[1:]) || brute(target, add, entriesLeft[1:])
}

func part2Func(input string) any {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		cases := strings.Split(line, ":")
		left := common.AtoI(cases[0])
		right := strings.Split(strings.TrimSpace(cases[1]), " ")

		if bruteStr(left, right[0], right[1:]) {
			total += left
		}
	}

	return total
}

func bruteStr(target int, total string, entriesLeft []string) bool {
	intTotal := common.AtoI(total)
	if len(entriesLeft) == 0 {
		return target == intTotal
	}

	nextInt := common.AtoI(entriesLeft[0])

	mul := intTotal * nextInt
	add := intTotal + nextInt

	mulStr := strconv.Itoa(mul)
	addStr := strconv.Itoa(add)
	conStr := total + entriesLeft[0]

	return bruteStr(target, mulStr, entriesLeft[1:]) || bruteStr(target, addStr, entriesLeft[1:]) || bruteStr(target, conStr, entriesLeft[1:])
}
