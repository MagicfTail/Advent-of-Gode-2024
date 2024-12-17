package main

import (
	"aoc-in-go/common"
	"regexp"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const inputRegex = `p=(\d*),(\d*) v=(-?\d*),(-?\d*)`

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

type entry struct {
	StartX int
	StartY int
	VelX int
	VelY int
}

func parseInput(input string) []entry {
	re := regexp.MustCompile(inputRegex)
	found := re.FindAllStringSubmatch(input, -1)

	entries := make([]entry, 0, len(found))

	for _, example := range found {
		entries = append(entries, entry{
			StartX: common.AtoI(example[1]),
			StartY: common.AtoI(example[2]),
			VelX:   common.AtoI(example[3]),
			VelY:   common.AtoI(example[4]),
		})
	}

	return entries
}

func part1Func(input string) any {
	entires := parseInput(input)

	boardWidth := 101 // 11 for example
	boardHeight := 103 // 7 for example

	finalPositions := make(map[common.Coordinate]int)

	for _, entry := range entires {
		finalX := (100*entry.VelX + entry.StartX) % boardWidth
		finalY := (100*entry.VelY + entry.StartY) % boardHeight

		finalPositions[common.Coordinate{X: finalX, Y: finalY}]++
	}

	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0

	middleRow := boardHeight / 2
	middleCol := boardWidth / 2

	for position, count := range finalPositions {
		if position.X < 0 {
			position.X += boardWidth
		}
		if position.Y < 0 {
			position.Y += boardHeight
		}

		switch {
		case position.X > middleCol && position.Y < middleRow:
			quad1 += count
		case position.X > middleCol && position.Y > middleRow:
			quad2 += count
		case position.X < middleCol && position.Y > middleRow:
			quad3 += count
		case position.X < middleCol && position.Y < middleRow:
			quad4 += count
		}
	}

	return quad1 * quad2 * quad3 * quad4
}

func makeClean(width, height int) [][]rune {
	out := make([][]rune, height)
	for i := range out {
		out[i] = make([]rune, width)
		for j := range out[i] {
			out[i][j] = '.'
		}
	}

	return out
}

func checkClean(in [][]rune) bool {
	for _, line := range in {
		if strings.Contains(string(line), "XXXXXXXXXX") {
			return true
		}
	}

	return false
}

func part2Func(input string) any {
	// Only works for user data
	entires := parseInput(input)

	boardWidth := 101 // 11 for example
	boardHeight := 103 // 7 for example

	count := 0

	for {
		count++
		out := makeClean(boardWidth, boardHeight)
		for i := range entires {
			entires[i].StartX = (entires[i].StartX + entires[i].VelX) % boardWidth
			entires[i].StartY = (entires[i].StartY + entires[i].VelY) % boardHeight

			if entires[i].StartX < 0 {
				entires[i].StartX += boardWidth
			}
			if entires[i].StartY < 0 {
				entires[i].StartY += boardHeight
			}

			out[entires[i].StartY][entires[i].StartX] = 'X'
		}
		if checkClean(out) {
			return count
		}
	}
}
