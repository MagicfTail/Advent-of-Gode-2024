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

func part1Func(input string) any {
	board := common.Parse2dArrayRune(input)

	startPos := common.Coordinate{}

	width := len(board[0])
	height := len(board)

	for y, line := range board {
		for x, char := range line {
			switch char {
			case 'S':
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	horCheats, verCheats := cheatPositions(board, height, width)

	stepsMap, _ := pathDistances(board, startPos)

	total := 0
	for _, horCheat := range horCheats {
		posA := common.Coordinate{X: horCheat.X - 1, Y: horCheat.Y}
		posB := common.Coordinate{X: horCheat.X + 1, Y: horCheat.Y}
		
		if common.IntAbs(stepsMap[posA] - stepsMap[posB]) - 2 >= 100 {
			total++
		}
	}
	for _, verCheat := range verCheats {
		posA := common.Coordinate{X: verCheat.X, Y: verCheat.Y - 1}
		posB := common.Coordinate{X: verCheat.X, Y: verCheat.Y + 1}
		
		if common.IntAbs(stepsMap[posA] - stepsMap[posB]) - 2 >= 100 {
			total++
		}
	}

	return total
}

func isCheat(a, b rune) bool {
	return (a == '.' || a == 'S' || a == 'E') && (b == '.' || b == 'S' || b == 'E')
}

func cheatPositions(board [][]rune, width, height int) ([]common.Coordinate, []common.Coordinate) {
	outHor := make([]common.Coordinate, 0)
	outVer := make([]common.Coordinate, 0)

	for y := range height - 2 {
		for x := range width - 2 {
			if board[y+1][x+1] != '#' {
				continue
			}

			if isCheat(board[y+1][x], board[y+1][x+2]) {
				outHor = append(outHor, common.Coordinate{X: x+1, Y: y+1})
			} else if isCheat(board[y][x+1], board[y+2][x+1]) {
				outVer = append(outVer, common.Coordinate{X: x+1, Y: y+1})
			}
		}
	}

	return outHor, outVer
}

func pathDistances(bord [][]rune, startPos common.Coordinate) (map[common.Coordinate]int, map[common.Coordinate]common.Direction) {
	pos := startPos

	stepsMap := make(map[common.Coordinate]int)
	dirMap := make(map[common.Coordinate]common.Direction)

	i := 0
	for {
		stepsMap[pos] = i
		if bord[pos.Y][pos.X] == 'E' {
			return stepsMap, dirMap
		}
		for _, direction := range common.CardinalDirections {
			newPos := common.Move(pos, direction)

			if _, ok := stepsMap[newPos]; ok {
				continue
			}

			if bord[newPos.Y][newPos.X] == '.' || bord[newPos.Y][newPos.X] == 'E' {
				dirMap[pos] = direction
				pos = newPos
				break
			}
		}
		i++
	}
}

func part2Func(input string) any {
	board := common.Parse2dArrayRune(input)

	startPos := common.Coordinate{}

	width := len(board[0])
	height := len(board)

	for y, line := range board {
		for x, char := range line {
			switch char {
			case 'S':
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	stepsMap, _ := pathDistances(board, startPos)

	savedMap := make(map[common.Coordinate]map[common.Coordinate]int)

	for y := range height {
		for x := range width {
			if board[y][x] == '#' {
				continue
			}

			pos := common.Coordinate{X: x, Y: y}

			checkAllSquares(savedMap, stepsMap, pos, width, height)
		}
	}

	total := 0
	
	for _, vals := range savedMap {
		for _, picoSaves := range vals {
			if picoSaves >= 100 {
				total++
			}
		}
	}

	return total/2
}

func checkAllSquares(savedMap map[common.Coordinate]map[common.Coordinate]int, stepsMap map[common.Coordinate]int, pos common.Coordinate, width, height int) {
	posSteps := stepsMap[pos]

	for tmpX := range 21 {
		for tmpY := range 21 - tmpX {
			for _, signX := range []int{-1, 1} {
				for _, signY := range []int{-1, 1} {
					y := tmpY * signY
					x := tmpX * signX
					otherPos := common.Coordinate{X: pos.X + x, Y: pos.Y + y}
					if common.IsOOB(otherPos.X, otherPos.Y, width, height) {
						continue
					}
					otherPosSteps, ok := stepsMap[otherPos]
					if !ok {
						continue
					}
					
					if common.InDoubleMap(savedMap, pos, otherPos) {
						continue
					}

					picoSaved := common.IntAbs(posSteps - otherPosSteps) - (common.IntAbs(x) + common.IntAbs(y))
					
					common.AddToDoubleMap(savedMap, pos, otherPos, picoSaved)
					common.AddToDoubleMap(savedMap, otherPos, pos, picoSaved)
				}
			}
		}
	}
}
