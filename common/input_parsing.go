package common

import (
	"strconv"
	"strings"
)

func IsOOB(x, y, width, height int) bool {
	if x < 0 || x >= width || y < 0 || y >= height {
		return true
	}

	return false
}

func ParseDoubleListInt(input string) (left, right []int) {
	lines := strings.Split(input, "\n")
	
	left = make([]int, 0, len(lines))
	right = make([]int, 0, len(lines))

	for _, line := range lines {
		values := strings.Fields(line)
		leftVal, err := strconv.Atoi(values[0])
		if err != nil {
			panic("Invalid parsing")
		}
		left = append(left, leftVal)

		rightVal, err := strconv.Atoi(values[1])
		if err != nil {
			panic("Invalid parsing")
		}
		right = append(right, rightVal)
	}

	return left, right
}

func Parse2dArrayIntSplit(input string) ([][]int) {
	lines := strings.Split(input, "\n")

	out := make([][]int, len(lines))

	for i, line := range lines {
		values := strings.Fields(line)
		for _, value := range values {
			val, err := strconv.Atoi(value)
			if err != nil {
				panic("Invalid parsing")
			}

			out[i] = append(out[i], val)
		}
	}

	return out
}

func Parse2dArrayRune(input string) ([][]rune) {
	lines := strings.Split(input, "\n")

	out := make([][]rune, len(lines))

	for i, line := range lines {
		for _, value := range line {
			out[i] = append(out[i], value)
		}
	}

	return out
}

func Parse2dArrayInt(input string) ([][]int) {
	lines := strings.Split(input, "\n")

	out := make([][]int, len(lines))

	for i, line := range lines {
		for _, value := range line {
			val, err := strconv.Atoi(string(value))
			if err != nil {
				panic("Invalid parsing")
			}

			out[i] = append(out[i], val)
		}
	}

	return out
}

func ParseListInt(input string) ([]int) {
	entries :=  strings.Split(input, " ")

	out := make([]int, 0, len(entries))
	for _, entry := range entries {
		out = append(out, AtoI(entry))
	}

	return out
}

func ParseCoordList(input string) []Coordinate {
	lines := strings.Split(input, "\n")

	coords := make([]Coordinate, 0, len(lines))

	for _, line := range lines {
		matches := strings.Split(line, ",")

		coords = append(coords, Coordinate{X: AtoI(matches[0]), Y: AtoI(matches[1])})
	}

	return coords
}
