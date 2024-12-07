package common

import (
	"fmt"
)

func ZipMap[T, G any](left, right []T, mapFunc func(a, b T)G) ([]G, error) {
	if len(left) != len(right) {
		return nil, fmt.Errorf("input length missmatch: left = %v, right = %v", len(left), len(right))
	}

	out := make([]G, len(left))

	for i := range out {
		out[i] = mapFunc(left[i], right[i])
	}

	return out, nil
}

func MapIgnoreErr[T, G any](in []T, mapFunc func(a T)(G, error)) ([]G) {
	out := make([]G, 0, len(in))

	for _, element := range in {
		value, _ := mapFunc(element)
		out = append(out, value)
	}

	return out
}

func Reduce[G any](in []G, initial G, reduceFunc func(a, b G) G) G {
	out := initial

	for _, val := range in {
		out = reduceFunc(out, val)
	}

	return out
}

func SliceToCountMap[T comparable](in []T) map[T]int {
	out := make(map[T]int)

	for _, entity := range in {
		out[entity] += 1
	}

	return out
}

func SliceToSet[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{})
	
	for _, entity := range in {
		out[entity] = struct{}{}
	}

	return out
}

func InterMap[T any](in []T, mapFunc func(T, T)T) []T {
	out := make([]T, 0, len(in) - 1)
	for i := 0; i < len(in) - 1; i++ {
		out = append(out, mapFunc(in[i], in[i + 1]))
	}

	return out
}

func SumBoolMatrix(in [][]bool) int {
	total := 0
	for _, inner := range in {
		for _, val := range inner {
			if val {
				total++
			}
		}
	}

	return total
}
