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
	// solve part 1 here
	return part1Func(input)
}

func part1Func(input string) any {
	mapData := common.Parse2dArrayRune(input)

	height := len(mapData)
	width := len(mapData[0])

	startX, startY := 0, 0
	out:
	for y := range height {
		for x := range width {
			if mapData[y][x] == '^' {
				startX, startY = x, y
				break out
			}
		}
	}

	seenSquares := make([][]bool, height)
	for i := range height {
		seenSquares[i] = make([]bool, width)
	}

	seenSquares[startY][startX] = true

	walkDirection(mapData, seenSquares, startX, startY, width, height, 0)

	return common.SumBoolMatrix(seenSquares)
}

func walkDirection(mapData [][]rune, seenSquares [][]bool, startX, startY, width, height, directionIndex int) {
	directionX := common.CardinalDirections[directionIndex].Horizontal
	directionY := common.CardinalDirections[directionIndex].Vertical
	x := startX + directionX
	y := startY + directionY

	for x >= 0 && x < width && y >=0 && y < height {
		if mapData[y][x] == '#' {
			walkDirection(mapData, seenSquares, x - directionX, y - directionY, width, height, int(math.Mod(float64(directionIndex + 1), 4)))
			return
		}
		seenSquares[y][x] = true


		x += common.CardinalDirections[directionIndex].Horizontal
		y += common.CardinalDirections[directionIndex].Vertical
	}
}

type spotWithDirection struct {
	X int
	Y int
	Direction int
}

func part2Func(input string) any {
	mapData := common.Parse2dArrayRune(input)

	height := len(mapData)
	width := len(mapData[0])

	startX, startY := 0, 0
	out:
	for y := range height {
		for x := range width {
			if mapData[y][x] == '^' {
				startX, startY = x, y
				break out
			}
		}
	}

	seenSquares := make([][]bool, height)
	for i := range height {
		seenSquares[i] = make([]bool, width)
	}

	seenSquares[startY][startX] = true

	walkDirection(mapData, seenSquares, startX, startY, width, height, 0)

	total := 0

	for i, inner := range seenSquares {
		for j, val := range inner {
			if val {
				seenSquaresTmp := make(map[spotWithDirection]bool)
				mapData[i][j] = '#'
				total += walkDirection2(mapData, seenSquaresTmp, startX, startY, width, height, 0)
				mapData[i][j] = '.'
			}
		}
	}

	return total
}

func walkDirection2(mapData [][]rune, seenSquares map[spotWithDirection]bool, startX, startY, width, height, directionIndex int) int {
	directionX := common.CardinalDirections[directionIndex].Horizontal
	directionY := common.CardinalDirections[directionIndex].Vertical
	x := startX
	y := startY

	for x >= 0 && x < width && y >=0 && y < height {
		if mapData[y][x] == '#' {
			return walkDirection2(mapData, seenSquares, x - directionX, y - directionY, width, height, int(math.Mod(float64(directionIndex + 1), 4)))
		}

		tmp := spotWithDirection{x, y, directionIndex}

		if seenSquares[tmp] {
			return 1
		} 

		seenSquares[tmp] = true


		x += common.CardinalDirections[directionIndex].Horizontal
		y += common.CardinalDirections[directionIndex].Vertical
	}

	return 0
}
