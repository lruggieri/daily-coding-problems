package util

func SliceSameUnordered[T comparable](container, elements []T) bool {
	if len(container) != len(elements) {
		return false
	}

	elementsMatched := map[T]struct{}{}

	for _, e1 := range container {
		var found bool
		for _, e2 := range elements {
			if _, alreadyMatched := elementsMatched[e2]; !alreadyMatched && e1 == e2 {
				elementsMatched[e2] = struct{}{}
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// MatrixSameUnordered : returns true if actual contains all elements of expected, even if in different order.
// Each raw of the matrix is matched using SliceSameUnordered.
func MatrixSameUnordered[T comparable](actual, expected [][]T) bool {
	if len(actual) != len(expected) {
		return false
	}

	elementsMatched := map[int]struct{}{}

	for _, e1 := range actual {
		var found bool
		for idx, e2 := range expected {
			if _, alreadyMatched := elementsMatched[idx]; !alreadyMatched && SliceSameUnordered(e1, e2) {
				elementsMatched[idx] = struct{}{}
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func Contains[T comparable](container []T, element T) bool {
	for _, e := range container {
		if e == element {
			return true
		}
	}

	return false
}
