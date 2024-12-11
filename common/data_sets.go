package common

type Direction struct {
	Vertical int
	Horizontal int
}

type Coordinate struct {
	X int
	Y int
}

var DigitSet = map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 0: {}}
var DigitSetString = map[string]struct{}{"1": {}, "2": {}, "3": {}, "4": {}, "5": {}, "6": {}, "7": {}, "8": {}, "9": {}, "0": {}}
var DigitSetRune = map[rune]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}, '0': {}}

var DirectionList = []int{-1, 0, 1}

// UP -> RIGHT -> DOWN -> LEFT
var CardinalDirections = []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
