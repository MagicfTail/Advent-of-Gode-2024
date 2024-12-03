package common

import "sort"

func SortIntSliceLt(in []int) {
	sort.Slice(in, func(i, j int) bool {
		return in[i] < in[j]
	})
}