package main

import (
	"aoc-in-go/common"
	"math"

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
	board := common.Parse2dArrayRune(input)

	minDistanceMap := make(map[common.Coordinate]map[common.Direction]int)

	endPos := common.Coordinate{}

	for y, line := range board {
		for x, char := range line {
			if char == 'S' {
				pos := common.Coordinate{X: x, Y: y}
				deptSearch(board, minDistanceMap, common.Right, pos, 0)
			}
			if char == 'E' {
				endPos = common.Coordinate{X: x, Y: y}
			}
		}
	}

	total := math.MaxInt
	for _, dist := range minDistanceMap[endPos] {
		total = common.IntMin(total, dist)
	}

	return total
}

func deptSearch(board [][]rune, minDistanceMap map[common.Coordinate]map[common.Direction]int, dir common.Direction, pos common.Coordinate, cost int) {
	if _, ok := minDistanceMap[pos]; !ok {
		minDistanceMap[pos] = make(map[common.Direction]int, 4)
	}
	if dist, ok := minDistanceMap[pos][dir]; ok && cost >= dist {
		return
	}
	minDistanceMap[pos][dir] = cost

	if board[pos.Y][pos.X] == 'E' {
		return
	}

	for _, direction := range common.CardinalDirections {
		if dir == common.InverDirection[direction] {
			continue
		}

		newPos := common.Move(pos, direction)
		if board[newPos.Y][newPos.X] != '#' {
			newCost := cost
			if dir == direction {
				newCost++
			} else {
				newCost += 1001
			}
			deptSearch(board, minDistanceMap, direction, newPos, newCost)
		}
	}
}

func part2Func(input string) any {
	board := common.Parse2dArrayRune(input)

	minDistanceMap := make(map[common.Coordinate]map[common.Direction]int)

	endPos := common.Coordinate{}

	for y, line := range board {
		for x, char := range line {
			if char == 'S' {
				pos := common.Coordinate{X: x, Y: y}
				deptSearch(board, minDistanceMap, common.Right, pos, 0)
			}
			if char == 'E' {
				endPos = common.Coordinate{X: x, Y: y}
			}
		}
	}

	added := make(map[common.Coordinate]bool)

	min := math.MaxInt
	for _, dist := range minDistanceMap[endPos] {
		min = common.IntMin(min, dist)
	}

	reverseDeptCount(board, minDistanceMap, added, endPos, min)

	for pos := range added {
		board[pos.Y][pos.X] = 'O'
	}

	return len(added)
}

func reverseDeptCount(board [][]rune, minDistanceMap map[common.Coordinate]map[common.Direction]int, added map[common.Coordinate]bool, pos common.Coordinate, cost int) {
	added[pos] = true

	if board[pos.Y][pos.X] == 'S' {
		return
	}

	for direction := range minDistanceMap[pos] {
		invDirection := common.InverDirection[direction]
		prevPos := common.Move(pos, invDirection)

		prevDirs := minDistanceMap[prevPos]
		for prevDir, prevDirDist := range prevDirs {
			if prevDir == direction && prevDirDist == cost - 1 {
				reverseDeptCount(board, minDistanceMap, added, prevPos, cost - 1)
			} else if prevDir != direction && prevDirDist == cost - 1001 {
				reverseDeptCount(board, minDistanceMap, added, prevPos, cost - 1001)
			}
		}
	}
}
