package main

import (
	"aoc-in-go/common"
	"fmt"

	"github.com/jpillora/puzzler/harness/aoc"
)

const isExample = false

const exampleBoardSize = 7
const exampleFallenBytes = 12
const userBoardSize = 71
const userFallenBytes = 1024

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
	coords := common.ParseCoordList(input)

	boardSize := common.Ternary(isExample, exampleBoardSize, userBoardSize)
	fallenBytes := common.Ternary(isExample, exampleFallenBytes, userFallenBytes)

	startCoord := common.Coordinate{X: 0, Y: 0}
	endCoord := common.Coordinate{X: boardSize - 1, Y: boardSize - 1}

	obstacleMap := make(map[common.Coordinate]bool, boardSize * boardSize)
	checkedMap := make(map[common.Coordinate]bool, boardSize * boardSize)
	for i := range fallenBytes {
		obstacleMap[coords[i]] = true
	}

	borderMap := map[common.Coordinate]struct{}{startCoord: {}}

	i := 0
	for {
		if len(borderMap) == 0 {
			return -1
		}
		newBorderMap := make(map[common.Coordinate]struct{})

		for borderCoord := range borderMap {
			if borderCoord == endCoord {
				return i
			}

			checkedMap[borderCoord] = true

			for _, direction := range common.CardinalDirections {
				checkCoord := common.Coordinate{X: borderCoord.X + direction.Horizontal, Y: borderCoord.Y + direction.Vertical}

				if common.IsOOB(checkCoord.X, checkCoord.Y, boardSize, boardSize) {
					continue
				}

				if obstacleMap[checkCoord] || checkedMap[checkCoord] {
					continue
				}

				newBorderMap[checkCoord] = struct{}{}
			}
		}

		borderMap = newBorderMap
		i++
	}
}

func part2Func(input string) any {
	coords := common.ParseCoordList(input)

	boardSize := common.Ternary(isExample, exampleBoardSize, userBoardSize)

	startCoord := common.Coordinate{X: 0, Y: 0}
	endCoord := common.Coordinate{X: boardSize - 1, Y: boardSize - 1}
	obstacleMap := make(map[common.Coordinate]bool, boardSize * boardSize)
	for _, coord := range coords {
		obstacleMap[coord] = true
		if !floodSearch(obstacleMap, startCoord, endCoord, boardSize) {
			return fmt.Sprintf("%v,%v", coord.X, coord.Y)
		}
	}

	return "-1,-1"
}

func floodSearch(obstacleMap map[common.Coordinate]bool, startCoord, endCoord common.Coordinate, boardSize int) bool {
	checkedMap := make(map[common.Coordinate]bool, boardSize * boardSize)
	borderMap := map[common.Coordinate]struct{}{startCoord: {}}

	for {
		if len(borderMap) == 0 {
			return false
		}
		newBorderMap := make(map[common.Coordinate]struct{})

		for borderCoord := range borderMap {
			if borderCoord == endCoord {
				return true
			}

			checkedMap[borderCoord] = true

			for _, direction := range common.CardinalDirections {
				checkCoord := common.Coordinate{X: borderCoord.X + direction.Horizontal, Y: borderCoord.Y + direction.Vertical}

				if common.IsOOB(checkCoord.X, checkCoord.Y, boardSize, boardSize) {
					continue
				}

				if obstacleMap[checkCoord] || checkedMap[checkCoord] {
					continue
				}

				newBorderMap[checkCoord] = struct{}{}
			}
		}

		borderMap = newBorderMap
	}
}
