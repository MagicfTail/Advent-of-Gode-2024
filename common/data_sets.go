package common


var Up = Direction{-1, 0}
var Down = Direction{1, 0}
var Left = Direction{0, -1}
var Right = Direction{0, 1}

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
var CardinalDirections = []Direction{Up, Right, Down, Left}

// UP-RIGHT -> DOWN-RIGHT -> DOWN-LEFT -> UP-LEFT
var XDirections = []Direction{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}

var UpDown = Set[Direction]{Up: {}, Down: {}}
var LeftRight = Set[Direction]{Left: {}, Right: {}}

var DirectionLookup = map[rune]Direction{'^': Up, '>': Right, 'v': Down, '<': Left}
var InverDirection = map[Direction]Direction{Up: Down, Right: Left, Down: Up, Left: Right}
