package main

import (
	"aoc-in-go/common"
	"os"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	b, _ := os.ReadFile("./input-user.txt")

	str := string(b)
	part1Func(str)
	part2Func(str)


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
	board := common.Parse2dArrayInt(input)

	pathCount := make(map[common.Coordinate]common.Set[common.Coordinate])

	height := len(board)
	width := len(board[0])

	total := 0

	for y, line := range board {
		for x, val := range line {
			if val != 0 {
				continue
			}
			found := find(board, pathCount, x, y, height, width, 0)
			total += len(found)
		}
	}
	
	return total
}

func find(board [][]int, pathMap map[common.Coordinate]common.Set[common.Coordinate], startX, startY, height, width, val int) common.Set[common.Coordinate] {
	if val == 9 {
		return map[common.Coordinate]struct{}{{X: startX, Y: startY}: {}}
	}

	coordinate := common.Coordinate{X: startX, Y: startY}

	if paths, seen := pathMap[coordinate]; seen {
		return paths
	}

	total := make(common.Set[common.Coordinate])

	for _, direction := range common.CardinalDirections {
		checkX := direction.Horizontal + startX
		checkY := direction.Vertical + startY

		if checkX < 0 || checkX >= width || checkY < 0 || checkY >= height {
			continue
		}

		nextVal := val + 1

		if board[checkY][checkX] != nextVal {
			continue
		}
		
		total = common.MergeSets(total, find(board, pathMap, checkX, checkY, height, width, nextVal))
	}

	pathMap[coordinate] = total
	
	return total
}

func part2Func(input string) any {
	board := common.Parse2dArrayInt(input)

	pathCount := make(map[common.Coordinate]int)

	height := len(board)
	width := len(board[0])

	total := 0

	for y, line := range board {
		for x, val := range line {
			if val != 0 {
				continue
			}
			total += findNumberOfRoutes(board, pathCount, x, y, height, width, 0)
		}
	}
	
	return total
}

func findNumberOfRoutes(board [][]int, pathMap map[common.Coordinate]int, startX, startY, height, width, val int) int {
	if val == 9 {
		return 1
	}

	coordinate := common.Coordinate{X: startX, Y: startY}

	if paths, seen := pathMap[coordinate]; seen {
		return paths
	}

	total := 0

	for _, direction := range common.CardinalDirections {
		checkX := direction.Horizontal + startX
		checkY := direction.Vertical + startY

		if checkX < 0 || checkX >= width || checkY < 0 || checkY >= height {
			continue
		}

		nextVal := val + 1

		if board[checkY][checkX] != nextVal {
			continue
		}
		
		total += findNumberOfRoutes(board, pathMap, checkX, checkY, height, width, nextVal)
	}

	pathMap[coordinate] = total
	
	return total
}
