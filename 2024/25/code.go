package main

import (
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var entryRegexExp = `.{5}
.{5}
.{5}
.{5}
.{5}
.{5}
.{5}`

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

type entry struct {
	first int
	second int
	third int
	fourth int
	fifth int
}

func parseInput(input string) ([]entry, []entry) {
	locks := make([]entry, 0)
	keys := make([]entry, 0)

	entryRegex := regexp.MustCompile(entryRegexExp)

	matches := entryRegex.FindAllString(input, -1)

	for _, match := range matches {
		lines := strings.Split(match, "\n")

		tmpEntry := make([]int, 5)
		for _, line := range lines {
			for i, char := range line {
				if char == '#' {
					tmpEntry[i]++
				}
			}
		}
		currentEntry := entry{
			first:  tmpEntry[0]-1,
			second: tmpEntry[1]-1,
			third:  tmpEntry[2]-1,
			fourth: tmpEntry[3]-1,
			fifth:  tmpEntry[4]-1,
		}

		if lines[0] == "#####" {
			locks = append(locks, currentEntry)
		} else {
			keys = append(keys, currentEntry)
		}
	}

	return locks, keys
}

func part1Func(input string) any {
	locks, keys := parseInput(input)

	total := 0

	for _, key := range keys {
		for _, lock := range locks {
			if key.first + lock.first > 5 || key.second + lock.second > 5 || key.third + lock.third > 5 || key.fourth + lock.fourth > 5 || key.fifth + lock.fifth > 5 {
				continue
			}

			total++
		}
	}

	return total
}

func part2Func(_ string) any {
	return "Merry Christmas!"
}
