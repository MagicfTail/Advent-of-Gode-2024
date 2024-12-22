package main

import (
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

func parseInput(input string) (towels, patterns []string) {
	lines := strings.Split(input, "\n")

	towels = strings.Split(lines[0], ", ")
	patterns = lines[2:]
	return
}

func part1Func(input string) any {
	towels, patterns := parseInput(input)

	towelsMap := make(map[string]bool)
	towelStripes := make(map[int]struct{})
	for _, towel := range towels {
		towelsMap[towel] = true
		towelStripes[len(towel)] = struct{}{}
	}

	seenPatterns := make(map[string]bool)

	total := 0
	for _, pattern := range patterns {
		if checkPattern(seenPatterns, pattern, towelsMap, towelStripes) {
			total++
		}
	}

	return total
}

func checkPattern(seenPatterns map[string]bool, pattern string, towelsMap map[string]bool, towelsStripes map[int]struct{}) bool {
	if possible, ok := seenPatterns[pattern]; ok {
		return possible
	}

	if pattern == "" {
		return true
	}

	for stripeCount := range towelsStripes {
		if len(pattern) < stripeCount {
			continue
		}
		towelPart := pattern[:stripeCount]

		if !towelsMap[towelPart] {
			continue
		}

		if checkPattern(seenPatterns, pattern[stripeCount:], towelsMap, towelsStripes) {
			seenPatterns[pattern] = true
			return true
		}
	}

	seenPatterns[pattern] = false
	return false
}

func part2Func(input string) any {
	towels, patterns := parseInput(input)

	towelsMap := make(map[string]bool)
	for _, towel := range towels {
		towelsMap[towel] = true
	}

	seenPatterns := make(map[string]int)

	total := 0
	for _, pattern := range patterns {
		total += countPattern(seenPatterns, pattern, towelsMap, towels)
	}

	return total
}

func countPattern(seenPatterns map[string]int, pattern string, towelsMap map[string]bool, towels []string) int {
	if possible, ok := seenPatterns[pattern]; ok {
		return possible
	}

	if pattern == "" {
		return 1
	}

	total := 0

	for _, towel := range towels {
		if !strings.HasPrefix(pattern, towel) {
			continue
		}

		total += countPattern(seenPatterns, pattern[len(towel):], towelsMap, towels)
	}

	seenPatterns[pattern] = total
	return total
}
