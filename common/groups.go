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

func SetEquals[T comparable](a, b Set[T]) bool {
	if len(a) != len(b) {
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

func SetToSlice[T comparable](set Set[T]) []T {
	out := make([]T, 0, len(set))
	for val := range set {
		out = append(out, val)
	}

	return out
}

func CopySet[T comparable](set Set[T]) Set[T] {
	newSet := make(Set[T], len(set))
	for key := range set {
		newSet[key] = struct{}{}
	}

	return newSet
}

func CopyMap[T comparable, G any](in map[T]G) map[T]G {
	newMap := make(map[T]G, len(in))
	for key, val := range in {
		newMap[key] = val
	}

	return newMap
}
