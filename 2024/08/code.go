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

func part1Func(input string) any {
	lines := strings.Split(input, "\n")

	antennas := make(map[rune][]common.Coordinate)
	
	width := len(lines)
	height := len(lines[0])

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}

			antennas[char] = append(antennas[char], common.Coordinate{
				X: x,
				Y: y,
			})
		} 
	}

	antinodes := make(map[common.Coordinate]struct{})
	
	for _, points := range antennas {
		if len(points) <= 1 {
			continue
		}

		for i, point1 := range points {
			for j, point2 := range points {
				if i == j {
					continue
				}

				pos := common.Coordinate{
					X: 2*point1.X - point2.X,
					Y: 2*point1.Y - point2.Y,
				}

				if !(pos.Y < 0 || pos.Y >= height || pos.X < 0 || pos.X >= width) {
					antinodes[pos] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func part2Func(input string) any {
	lines := strings.Split(input, "\n")

	antennas := make(map[rune][]common.Coordinate)
	
	width := len(lines)
	height := len(lines[0])

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}

			antennas[char] = append(antennas[char], common.Coordinate{
				X: x,
				Y: y,
			})
		} 
	}

	antinodes := make(map[common.Coordinate]struct{})
	
	for _, points := range antennas {
		if len(points) <= 1 {
			continue
		}

		for i, point1 := range points {
			for j, point2 := range points {
				if i == j {
					continue
				}

				dif := common.Coordinate{
					X: point1.X - point2.X,
					Y: point1.Y - point2.Y,
				}

				pos := common.Coordinate{
					X: point1.X,
					Y: point1.Y,
				}

				for !(pos.Y < 0 || pos.Y >= height || pos.X < 0 || pos.X >= width) {
					antinodes[pos] = struct{}{}

					pos = common.Coordinate{
					X: pos.X + dif.X,
					Y: pos.Y + dif.Y,
				}
				}
			}
		}
	}

	return len(antinodes)
}
