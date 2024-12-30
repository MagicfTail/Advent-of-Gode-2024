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

func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
} 

func AddToDoubleMap[T, G, H comparable](in map[T]map[G]H, t T, g G, h H) {
	if _, ok := in[t]; !ok {
		in[t] = make(map[G]H)
	}

	in[t][g] = h
}

func InDoubleMap[T, G, H comparable](in map[T]map[G]H, t T, g G) bool {
	inner, ok := in[t]
	if !ok {
		return false
	}

	_, ok = inner[g]
	return ok
}

func IsOOB(x, y, width, height int) bool {
	if x < 0 || x >= width || y < 0 || y >= height {
		return true
	}

	return false
}

func LPad(s string,pad string, plength int) string {
    for i:=len(s);i<plength;i++{
        s=pad+s
    }
    return s
}
