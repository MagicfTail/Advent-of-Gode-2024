package common

func IntDiff(left, right int) int {
	return IntAbs(IntSub(left, right))
}

func IntAdd(left, right int) int {
	return left + right
}

func IntSub(left, right int) int {
	return left - right
}

func IntMul(left, right int) int {
	return left * right
}

func IntSubReverse(left, right int) int {
	return right - left
}

func IntAbs(in int) int {
	if in < 0 {
		return -in
	} else {
		return in
	}
}

func IntSign(in int) int {
	if in < 0 {
		return -1
	} else {
		return 1
	}
}
