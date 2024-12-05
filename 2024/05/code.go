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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2Func(input)
	}
	// solve part 1 here
	return part1Func(input)
}

func part1Func(input string) int {
	lines := strings.Split(input, "\n")

	requiredPages := make(map[int][]int)
	i := 0
	for _, line := range lines {
		i++
		if len(line) == 0 {
			break
		}

		split := strings.Split(line, "|")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		requiredPages[left] =append(requiredPages[left], right)
	}

	total := 0

	for _, line := range lines[i:] {
		numbers := strings.Split(line, ",")
		ints := common.MapIgnoreErr(numbers, strconv.Atoi)

		total += checkValidOrder(requiredPages, ints)
	}

	return total	
}

func checkValidOrder(requiredPages map[int][]int, problem []int) int {
	seen := make(map[int]bool)

	for _, number := range problem {
			seen[number] = true
			required := requiredPages[number]

			for _, element := range required {
				if seen[element] {
					return 0
				}
			}
		
	}

	return problem[len(problem)/2]
}

func part2Func(input string) int {
	lines := strings.Split(input, "\n")

	requiredPages := make(map[int][]int)
	i := 0
	for _, line := range lines {
		i++
		if len(line) == 0 {
			break
		}

		split := strings.Split(line, "|")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		requiredPages[left] = append(requiredPages[left], right)
	}

	total := 0

	for _, line := range lines[i:] {
		numbers := strings.Split(line, ",")
		ints := common.MapIgnoreErr(numbers, strconv.Atoi)

		total += orderPages(requiredPages, ints)
	}

	return total	
}

func orderPages(requiredPages map[int][]int, problem []int) int {
	seen := make(map[int]bool)

	moved := false

	for i, number := range problem {
		seen[number] = true
		required := requiredPages[number]

		needPassing := make(map[int]bool)

		for _, element := range required {
			if seen[element] {
				needPassing[element] = true
			}
		}

		if len(needPassing) == 0 {
			continue
		}

		moved = true
		bubble(problem, needPassing, i)
	}

	if !moved {
		return 0
	}

	return problem[len(problem)/2]
}

func bubble(problem []int, needPassing map[int]bool, position int) {
	i := position
	for len(needPassing) > 0 {
		tmp := problem[i - 1]
		problem[i - 1] = problem[i]
		problem[i] = tmp
		
		i--
		delete(needPassing, tmp)
	}
}
