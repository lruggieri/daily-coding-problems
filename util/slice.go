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
