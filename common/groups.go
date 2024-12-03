package common

func Subset[T comparable](a, b map[T]struct{}) bool {
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