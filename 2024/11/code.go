package main

import (
	"aoc-in-go/common"
	"strconv"

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
	numbers := common.ParseListInt(input)

	for range 25 {
		tmp := make([]int, 0, len(numbers))
		for _, number := range numbers {
			if number == 0 {
				tmp = append(tmp, 1)
				continue
			}

			strNumber := strconv.Itoa(number)
			if len(strNumber)%2 == 0 {
				left := strNumber[:len(strNumber)/2]
				right := strNumber[len(strNumber)/2:]

				tmp = append(tmp, common.AtoI(left), common.AtoI(right))
				continue
			}

			tmp = append(tmp, number*2024)
		}
		numbers = tmp
	}

	return len(numbers)
}

type numberDepth struct {
	StartNumber int
	Depth       int
}

func part2Func(input string) any {
	numbers := common.ParseListInt(input)

	seen := make(map[numberDepth]int)

	total := 0

	for _, number := range numbers {
		total += Dive(seen, number, 0)
	}

	return total
}

func Dive(seen map[numberDepth]int, number, depth int) int {
	if depth == 75 {
		return 1
	}

	entry := numberDepth{StartNumber: number, Depth: depth}
	if val, ok := seen[entry]; ok {
		return val
	}

	if number == 0 {
		tmp := Dive(seen, 1, depth+1)
		seen[entry] = tmp
		return tmp
	}

	strNumber := strconv.Itoa(number)
	if len(strNumber)%2 == 0 {
		left := common.AtoI(strNumber[:len(strNumber)/2])
		right := common.AtoI(strNumber[len(strNumber)/2:])

		tmp := Dive(seen, left, depth+1) + Dive(seen, right, depth+1)
		seen[entry] = tmp
		return tmp
	}

	tmp := Dive(seen, number*2024, depth+1)
	seen[entry] = tmp
	return tmp
}
