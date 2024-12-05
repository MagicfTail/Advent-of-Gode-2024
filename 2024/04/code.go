package main

import (
	"aoc-in-go/common"

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

var xmas = []rune{'X', 'M', 'A', 'S'}

func part1Func(input string) any {
	parsedInput := common.Parse2dArrayRune(input)

	height := len(parsedInput)
	width := len(parsedInput[0])

	total := 0

	for y := range height {
		for x := range width {
			if parsedInput[y][x] == 'X' {
				total += searchInAllDirections(parsedInput, x, y, height, width)
			}
		}
	}

	return total
}

func searchInAllDirections(board [][]rune, startX, startY, height, width int) int {
	total := 0
	for _, x := range common.DirectionList {
		for _, y := range common.DirectionList {
			if x == 0 && y == 0 {
				continue
			}

			total += searchInDirection(board, startX, startY, x, y, height, width, 1)
		}
	}

	return total
}

func searchInDirection(board [][]rune, startX, startY, directionX, directionY, height, width, charNumber int) int {
	if charNumber == len(xmas) {
		return 1
	}

	x := startX + directionX
	y := startY + directionY

	if x >=0 && x < width && y >=0 && y < height {
		if board[y][x] == xmas[charNumber] {
			return searchInDirection(board, x, y, directionX, directionY, height, width, charNumber + 1)
		}
		return 0
	}

	return 0
}

func part2Func(input string) any {
	parsedInput := common.Parse2dArrayRune(input)

	height := len(parsedInput)
	width := len(parsedInput[0])

	total := 0

	for y := range height {
		for x := range width {
			if parsedInput[y][x] == 'A' {
				total += searchInX(parsedInput, x, y, height, width)
			}
		}
	}

	return total
}

func searchInX(board [][]rune, startX, startY, height, width int) int {
	dig11 := checkPosition(board, startX + 1, startY + 1, height, width, 'M') && checkPosition(board, startX - 1, startY - 1, height, width, 'S')
	dig12 := checkPosition(board, startX + 1, startY + 1, height, width, 'S') && checkPosition(board, startX - 1, startY - 1, height, width, 'M')

	dig21 := checkPosition(board, startX + 1, startY - 1, height, width, 'M') && checkPosition(board, startX - 1, startY + 1, height, width, 'S')
	dig22 := checkPosition(board, startX + 1, startY - 1, height, width, 'S') && checkPosition(board, startX - 1, startY + 1, height, width, 'M')

	if (dig11 || dig12) && (dig21 || dig22) {
		return 1
	}

	return 0
}

func checkPosition(board [][]rune, x, y, height, width int, char rune) bool {
	if x >=0 && x < width && y >=0 && y < height {
		if board[y][x] == char {
			return true
		}
	}

	return false
}

// Overlap does apparently not mean that there's allowed to be holes between the letters in XMAS. -30 min
// func searchInDirection(board [][]rune, startX, startY, directionX, directionY, height, width, charNumber int) int {
// 	if charNumber == len(xmas) {
// 		return 1
// 	}

// 	x := startX + directionX
// 	y := startY + directionY

// 	for x >=0 && x < width && y >=0 && y < height {
// 		if board[y][x] == xmas[charNumber] {
// 			return searchInDirection(board, x, y, directionX, directionY, height, width, charNumber + 1)
// 		}
// 		x += directionX
// 		y += directionY
// 	}

// 	return 0
// }
