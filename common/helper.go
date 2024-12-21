package common

import (
	"fmt"
)

func SwapMatrix[T comparable](matrix [][]T, a, b Coordinate) {
	tmp := matrix[a.Y][a.X]
	matrix[a.Y][a.X] = matrix[b.Y][b.X]
	matrix[b.Y][b.X] = tmp
}

func PrintRuneMatrix(matrix [][]rune) {
	for _, line := range matrix {
		fmt.Println(string(line))
	}
}

func Move(pos Coordinate, dir Direction) Coordinate {
	return Coordinate{X: pos.X + dir.Horizontal, Y: pos.Y + dir.Vertical}
}
