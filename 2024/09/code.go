package main

import (
	"aoc-in-go/common"
	"container/heap"
	"math"
	"os"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	b, _ := os.ReadFile("./input-user.txt")
	str := string(b)
	part2Func(str)
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
	values := make([]int, 0, len(input)/2)
	spaces := make([]int, 0, len(input)/2)

	for i, char := range input {
		val := common.AtoI(string(char))
		if i%2 == 0 {
			values = append(values, val)
		} else {
			spaces = append(spaces, val)
		}
	}

	total := 0

	onValue := true
	spaceCount := spaces[0]
	leftId := 0
	rightId := len(values) - 1

	position := 0

	for {
		if leftId == rightId && values[leftId] == 0 {
			return total
		}

		if onValue && values[leftId] == 0 {
			onValue = !onValue
			leftId++
			continue
		}

		if !onValue && spaceCount == 0 {
			onValue = !onValue
			spaces = spaces[1:]
			spaceCount = spaces[0]
		}

		if onValue {
			total += position * leftId
			values[leftId]--
		} else {
			if values[rightId] == 0 {
				rightId--
				continue
			}

			
			total += position * rightId
			values[rightId]--
			spaceCount--
		}

		position++
	}
}

type value struct {
	StartPos int
	Value int
	Length int
}

func part2Func(input string) any {
	values := make([]value, 0, len(input)/2)
	spaces := make(map[int]*common.IntHeap)

	for i := 0; i <= 9; i++ {
		var iHeap common.IntHeap
		heap.Init(&iHeap)
		spaces[i] = &iHeap
	}

	position := 0
	for i, char := range input {
		val := common.AtoI(string(char))
		if i%2 == 0 {
			values = append(values, value{
				StartPos: position,
				Value:    i/2,
				Length:   val,
			})
		} else {
			heap.Push(spaces[val], position)
		}
		position += val
	}

	newValues := make([]value, 0, len(input)/2)

	for i := len(values) - 1; i >= 0; i-- {
		val := values[i]

		firstFit := math.MaxInt
		firstSpotSize := 0

		for spotSize := val.Length; spotSize <= 9; spotSize++ {
			if (*spaces[spotSize]).Len() == 0 {
				continue
			}

			firstSpotPos := spaces[spotSize].Peek().(int)
			
			if firstSpotPos >= val.StartPos {
				continue
			}

			if firstFit > firstSpotPos {
				firstFit = firstSpotPos
				firstSpotSize = spotSize
			}
		}

		if firstFit == math.MaxInt {
			newValues = append(newValues, val)
			continue
		}

		_ = heap.Pop(spaces[firstSpotSize])

		val.StartPos = firstFit
		newSpaceSize := firstSpotSize - val.Length
		newSpacePos := firstFit + val.Length
		
		heap.Push(spaces[newSpaceSize], newSpacePos)

		newValues = append(newValues, val)
	}

	total := 0

	for _, value := range newValues {
		for i := range value.Length {
			total += value.Value * (value.StartPos + i)
		}
	}

	return total
}
