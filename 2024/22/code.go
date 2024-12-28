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

func mix(in, secret int) int {
	return in ^ secret
}

func prune(in int) int {
	return in % 16777216
}

func part1Func(input string) any {
	startingNumbers := common.ParseListInt2(input)

	total := 0
	for _, initialSecret := range startingNumbers {
		secret := initialSecret
		for range 2000 {
			secret = mix(secret * 64, secret)
			secret = prune(secret)
			secret = mix(secret / 32, secret)
			secret = prune(secret)
			secret = mix(secret * 2048, secret)
			secret = prune(secret)
		}

		total += secret
	}

	return total
}

type sequence struct {
	first int
	second int
	third int
	fourth int
}

func part2Func(input string) any {
	startingNumbers := common.ParseListInt2(input)
	diffSequences := make([][]int, 0, len(startingNumbers))
	valueSlices := make([][]int, 0, len(startingNumbers))

	for _, initialSecret := range startingNumbers {
		secret := initialSecret
		valueSlice := make([]int, 0, 2001)
		valueSlice = append(valueSlice, secret % 10)
		for range 2000 {
			secret = mix(secret * 64, secret)
			secret = prune(secret)
			secret = mix(secret / 32, secret)
			secret = prune(secret)
			secret = mix(secret * 2048, secret)
			secret = prune(secret)
			valueSlice = append(valueSlice, secret % 10)
		}
		valueSlices = append(valueSlices, valueSlice)
		diffSequences = append(diffSequences, common.InterMap(valueSlice, common.IntSubReverse))
	}

	sequenceMaps := make([]map[sequence]int, 0, len(startingNumbers))
	addedSequences := make(map[sequence]struct{})

	for i, diffSequnce := range diffSequences {
		sequenceMap := make(map[sequence]int)
		for j := range len(diffSequnce) - 3 {
			currentSequence := sequence{diffSequnce[j], diffSequnce[j+1], diffSequnce[j+2], diffSequnce[j+3]}
			if _, ok := sequenceMap[currentSequence]; ok {
				continue
			}

			sequenceMap[currentSequence] = valueSlices[i][j+4]
			addedSequences[currentSequence] = struct{}{}
		}
		sequenceMaps = append(sequenceMaps, sequenceMap)
	}
	total := 0
	
	for currentSequence := range addedSequences {

		currentTotal := 0
		for _, sequenceMap := range sequenceMaps {
			currentTotal += sequenceMap[currentSequence]
		}

		total = common.IntMax(total, currentTotal)
	}		

	return total
}
