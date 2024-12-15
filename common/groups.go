package common

type Set[T comparable] map[T]struct{}

func Subset[T comparable](a, b Set[T]) bool {
	if len(a) > len(b) {
		return false
	}

	for key := range a {
		if _, ok := b[key]; !ok {
			return false
		}
	}

	return true
}

func MergeSets[T comparable](a, b Set[T]) Set[T] {
	for val := range b {
		a[val] = struct{}{}
	}

	return a
}
