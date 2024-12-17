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
	if part2 {
		return part2Func(input)
	}

	return part1Func(input)
}

type field struct {
	Perimiter int
	Area      int
}

func part1Func(input string) any {
	inputMap := common.Parse2dArrayRune(input)

	height := len(inputMap)
	width := len(inputMap[0])

	seen := make(map[common.Coordinate]bool)

	total := 0

	for y, line := range inputMap {
		for x, char := range line {
			currentField := field{0, 0}
			exploreField(inputMap, seen, &currentField, x, y, width, height, char)

			total += currentField.Area * currentField.Perimiter
		}
	}

	return total
}

func exploreField(fields [][]rune, seen map[common.Coordinate]bool, currentField *field, x, y, width, height int, fieldType rune) {
	cord := common.Coordinate{X: x, Y: y}
	if seen[cord] {
		return
	}

	seen[cord] = true

	currentField.Area++

	for _, d := range common.CardinalDirections {
		newX := x + d.Horizontal
		newY := y + d.Vertical

		if common.IsOOB(newX, newY, width, height) {
			currentField.Perimiter++
			continue
		}

		if fields[newY][newX] != fieldType {
			currentField.Perimiter++
			continue
		}

		exploreField(fields, seen, currentField, newX, newY, width, height, fieldType)
	}
}

func part2Func(input string) any {
	inputMap := common.Parse2dArrayRune(input)

	height := len(inputMap)
	width := len(inputMap[0])

	seen := make(map[common.Coordinate]bool)

	total := 0

	for y, line := range inputMap {
		for x, char := range line {
			cord := common.Coordinate{
				X: x,
				Y: y,
			}

			if seen[cord] {
				continue
			}

			currentField := field{Perimiter: 0, Area: 0}
			exploreField2(inputMap, seen, &currentField, x, y, width, height, char)
			
			total += currentField.Area * currentField.Perimiter
		}
	}

	return total
}

func exploreField2(fields [][]rune, seen map[common.Coordinate]bool, currentField *field, x, y, width, height int, fieldType rune) {
	cord := common.Coordinate{X: x, Y: y}
	if seen[cord] {
		return
	}

	seen[cord] = true

	currentField.Area++

	checkOuter(fields, currentField, x, y, width, height, fieldType)
	checkInner(fields, currentField, x, y, width, height, fieldType)

	for _, d := range common.CardinalDirections {
		newX := x + d.Horizontal
		newY := y + d.Vertical

		if common.IsOOB(newX, newY, width, height) {
			continue
		}

		if fields[newY][newX] != fieldType {
			continue
		}

		exploreField2(fields, seen, currentField, newX, newY, width, height, fieldType)
	}
}

func checkOuter(fields [][]rune, currentField *field, x, y, width, height int, fieldType rune) {
	matchCount := 0
	for _, direction := range common.XDirections {
		check1X := x + direction.Horizontal
		check1Y := y
		check2X := x
		check2Y := y + direction.Vertical

		if (common.IsOOB(check1X, check1Y, width, height) || fields[check1Y][check1X] != fieldType) && (common.IsOOB(check2X, check2Y, width, height) || fields[check2Y][check2X] != fieldType) {
			matchCount++
		}
	}

	currentField.Perimiter += matchCount
}

func checkInner(fields [][]rune, currentField *field, x, y, width, height int, fieldType rune) {
	matchCount := 0
	for _, direction := range common.XDirections {
		newX := x + direction.Horizontal
		newY := y + direction.Vertical

		if common.IsOOB(newX, newY, width, height) {
			continue
		}

		if fields[y][newX] == fieldType && fields[newY][x] == fieldType && fields[newY][newX] != fieldType {
			matchCount++
		}
	}

	currentField.Perimiter += matchCount
}
