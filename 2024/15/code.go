package main

import (
	"aoc-in-go/common"
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

func parseInput(input string) (string, [][]rune, common.Coordinate) {
	lines := strings.Split(input, "\n")

	board := make([][]rune, 0, len(lines))
	cord := common.Coordinate{}
	moves := ""

	moveMode := false

	for y, line := range lines {
		if moveMode {
			moves += line
			continue
		}

		if len(line) == 0 {
			moveMode = true
			continue
		}

		chars := make([]rune, 0, len(line))
		for x, char := range line {
			chars = append(chars, char)
			if char == '@' {
				cord = common.Coordinate{X: x, Y: y}
			}
		}
		board = append(board, chars)
	}

	return moves, board, cord
}

func part1Func(input string) any {
	moves, board, cord := parseInput(input)

	for _, move := range moves {
		direction := common.DirectionLookup[move]
		if deepSwap(board, direction, cord) {
			cord.X += direction.Horizontal
			cord.Y += direction.Vertical
		}
	}

	total := 0
	for y, line := range board {
		for x, char := range line {
			if char == 'O' {
				total += y * 100 + x
			}
		}
	}

	return total
}

func deepSwap(board [][]rune, direction common.Direction, pos common.Coordinate) bool {
	newPos := common.Coordinate{X: pos.X + direction.Horizontal, Y: pos.Y + direction.Vertical}

	if board[newPos.Y][newPos.X] == '#' {
		return false
	}
	if board[newPos.Y][newPos.X] == '.' {
		common.SwapMatrix(board, newPos, pos)
		return true
	}
	if deepSwap(board, direction, newPos) {
		common.SwapMatrix(board, newPos, pos)
		return true
	}

	return false
}

func parseInputWide(input string) (string, [][]rune, common.Coordinate) {
	lines := strings.Split(input, "\n")

	board := make([][]rune, 0, len(lines))
	cord := common.Coordinate{}
	moves := ""

	moveMode := false

	for y, line := range lines {
		if moveMode {
			moves += line
			continue
		}

		if len(line) == 0 {
			moveMode = true
			continue
		}

		chars := make([]rune, 0, len(line) * 2)
		for x, char := range line {
			switch char {
			case '#':
				chars = append(chars, '#', '#')
			case 'O':
				chars = append(chars, '[', ']')
			case '.':
				chars = append(chars, '.', '.')
			case '@':
				chars = append(chars, '@', '.')
				cord = common.Coordinate{X: x * 2, Y: y}
			}
		}
		board = append(board, chars)
	}

	return moves, board, cord
}

func part2Func(input string) any {
	moves, board, cord := parseInputWide(input)

	for _, move := range moves {
		direction := common.DirectionLookup[move]
		swapDepth := deepFindSwapDepth(board, direction, cord, 0)
		if swapDepth > 0 {
			deepSwapWide(board, direction, cord, swapDepth)
			cord.X += direction.Horizontal
			cord.Y += direction.Vertical
		}
	}

	total := 0
	for y, line := range board {
		for x, char := range line {
			if char == '[' {
				total += y * 100 + x
			}
		}
	}

	return total
}

func deepFindSwapDepth(board [][]rune, direction common.Direction, pos common.Coordinate, depth int) int {
	newPos := common.Coordinate{X: pos.X + direction.Horizontal, Y: pos.Y + direction.Vertical}

	if board[newPos.Y][newPos.X] == '#' {
		return depth
	}
	if board[newPos.Y][newPos.X] == '.' {
		return depth + 1
	}

	if direction == common.CardinalDirections[0] || direction == common.CardinalDirections[2] {
		if board[newPos.Y][newPos.X] == '[' {
			offsetPos := common.Coordinate{X: newPos.X + 1, Y: newPos.Y}
			leftMin := deepFindSwapDepth(board, direction, newPos, depth)
			rightMin := deepFindSwapDepth(board, direction, offsetPos, depth)
			return common.IntMin(leftMin, rightMin)
		}
		if board[newPos.Y][newPos.X] == ']' {
			offsetPos := common.Coordinate{X: newPos.X - 1, Y: newPos.Y}
			leftMin := deepFindSwapDepth(board, direction, offsetPos, depth)
			rightMin := deepFindSwapDepth(board, direction, newPos, depth)
			return common.IntMin(leftMin, rightMin)
		}
	}

	return deepFindSwapDepth(board, direction, newPos, depth)
}

func deepSwapWide(board [][]rune, direction common.Direction, pos common.Coordinate, depth int) {
	if depth == 0 {
		return
	}
	newPos := common.Coordinate{X: pos.X + direction.Horizontal, Y: pos.Y + direction.Vertical}

	if board[newPos.Y][newPos.X] == '.' {
		deepSwapWide(board, direction, newPos, depth - 1)
		common.SwapMatrix(board, newPos, pos)
		return
	}
	
	if direction == common.CardinalDirections[0] || direction == common.CardinalDirections[2] {
		if board[newPos.Y][newPos.X] == '[' {
			offsetPos := common.Coordinate{X: newPos.X + 1, Y: newPos.Y}
			deepSwapWide(board, direction, offsetPos, depth)
		}
		if board[newPos.Y][newPos.X] == ']' {
			offsetPos := common.Coordinate{X: newPos.X - 1, Y: newPos.Y}
			deepSwapWide(board, direction, offsetPos, depth)
		}
	}

	deepSwapWide(board, direction, newPos, depth)
	common.SwapMatrix(board, newPos, pos)
}
