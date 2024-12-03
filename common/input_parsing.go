package common

import (
	"strconv"
	"strings"
)

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

func Parse2dArrayInt(input string) ([][]int) {
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
