package main

import (
	"aoc-in-go/common"
	"os"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

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

func parseInput(input string) map[string]map[string]bool {
	lines := strings.Split(input, "\n")

	out := make(map[string]map[string]bool)

	for _, line := range lines {
		part1 := line[:2]
		part2 := line[3:]

		if _, ok := out[part1]; !ok {
			out[part1] = make(map[string]bool)
		}
		if _, ok := out[part2]; !ok {
			out[part2] = make(map[string]bool)
		}

		out[part1][part2] = true
		out[part2][part1] = true
	}

	return out
}

func part1Func(input string) any {
	connections := parseInput(input)

	total := 0
	for part1, part2s := range connections {
		for part2 := range part2s {
			for part3 := range connections[part2] {
				if connections[part1][part3] {
					if part1[0] == 't' || part2[0] == 't' || part3[0] == 't' {
						total++
					}
				}
			}
		}
	}

	return total/6
}

// Max cliques problem. NP-complete algorithm, solved using greedy algorithm. Is not guarenteed to be correct
func part2Func(input string) any {
	connections := parseInput(input)

	cliques := make([]map[string]struct{}, 0)

	for val1, val2s := range connections {
		out:
		for _, clique := range cliques {
			for existingEntries := range clique {
				if !val2s[existingEntries] {
					continue out
				}
			}
			clique[val1] = struct{}{}
		}
		cliques = append(cliques, map[string]struct{}{val1: {}})
	}

	maxLengthClique := map[string]struct{}{}
	for _, cluqiue := range cliques {
		if len(cluqiue) > len(maxLengthClique) {
			maxLengthClique = cluqiue
		}
	}

	cliqueSlice := common.SetToSlice(maxLengthClique)
	sort.Slice(cliqueSlice, func(i, j int) bool {
		return cliqueSlice[i] < cliqueSlice[j]
	})

	return strings.Join(cliqueSlice, ",")
}

