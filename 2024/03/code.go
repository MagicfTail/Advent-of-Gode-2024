package main

import (
	"aoc-in-go/common"
	"strconv"

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

var mulChars = []rune{'m', 'u', 'l', '(', ',', ')'}
var doChars = []rune{'d', 'o', '(', ')'}
var dontChars = []rune{'d', 'o', 'n', '\'', 't', '(', ')'}

func part1Func(input string) any {
	stringNumbersLeft := []string{}
	stringNumbersRight := []string{}

	currentIndex := 0
	currentNumberLeft := ""
	currentNumberRight := ""

	for _, letter := range input {
		if currentIndex == 4 {
			if _, ok := common.DigitSetRune[letter]; ok {
				currentNumberLeft += string(letter)
				continue
			}
		} else if currentIndex == 5 {
			if _, ok := common.DigitSetRune[letter]; ok {
				currentNumberRight += string(letter)
				continue
			}
		}

		if letter != mulChars[currentIndex] {
			currentIndex = 0
			currentNumberLeft = ""
			currentNumberRight = ""
			continue
		}

		if letter == ')' {
			stringNumbersLeft = append(stringNumbersLeft, currentNumberLeft)
			stringNumbersRight = append(stringNumbersRight, currentNumberRight)

			currentIndex = 0
			currentNumberLeft = ""
			currentNumberRight = ""
			continue
		}

		currentIndex += 1
	}

	intSliceLeft := make([]int, 0, len(stringNumbersLeft))
	for _, numberString := range stringNumbersLeft {
		number, _ := strconv.Atoi(numberString)
		intSliceLeft = append(intSliceLeft, number)
	}

	intSliceRight := make([]int, 0, len(stringNumbersRight))
	for _, numberString := range stringNumbersRight {
		number, _ := strconv.Atoi(numberString)
		intSliceRight = append(intSliceRight, number)
	}

	zipped, _ := common.ZipMap(intSliceLeft, intSliceRight, common.IntMul)
	
	return common.Reduce(zipped, 0, common.IntAdd)
}

func part2Func(input string) any {
	stringNumbersLeft := []string{}
	stringNumbersRight := []string{}

	currentIndexMul := 0
	currentIndexDo := 0
	currentIndexDont := 0
	currentNumberLeft := ""
	currentNumberRight := ""
	going := true

	for _, letter := range input {
		if letter != doChars[currentIndexDo] {
			currentIndexDo = -1
		}

		if letter != dontChars[currentIndexDont] {
			currentIndexDont = -1
		}

		if currentIndexDo == len(doChars) - 1 {
			going = true
			currentIndexDo = -1
		}

		if currentIndexDont == len(dontChars) - 1{
			going = false
			currentIndexDont = -1
		}

		currentIndexDont += 1
		currentIndexDo += 1

		if currentIndexMul == 4 {
			if _, ok := common.DigitSetRune[letter]; ok {
				currentNumberLeft += string(letter)
				continue
			}
		} else if currentIndexMul == 5 {
			if _, ok := common.DigitSetRune[letter]; ok {
				currentNumberRight += string(letter)
				continue
			}
		}

		if !going || letter != mulChars[currentIndexMul] {
			currentIndexMul = 0
			currentNumberLeft = ""
			currentNumberRight = ""
			continue
		}

		if letter == ')' && currentIndexMul == 5 {
			stringNumbersLeft = append(stringNumbersLeft, currentNumberLeft)
			stringNumbersRight = append(stringNumbersRight, currentNumberRight)

			currentIndexMul = 0
			currentNumberLeft = ""
			currentNumberRight = ""
			continue
		}

		currentIndexMul += 1
	}

	intSliceLeft := make([]int, 0, len(stringNumbersLeft))
	for _, numberString := range stringNumbersLeft {
		number, _ := strconv.Atoi(numberString)
		intSliceLeft = append(intSliceLeft, number)
	}

	intSliceRight := make([]int, 0, len(stringNumbersRight))
	for _, numberString := range stringNumbersRight {
		number, _ := strconv.Atoi(numberString)
		intSliceRight = append(intSliceRight, number)
	}

	zipped, _ := common.ZipMap(intSliceLeft, intSliceRight, common.IntMul)
	
	return common.Reduce(zipped, 0, common.IntAdd)
}
