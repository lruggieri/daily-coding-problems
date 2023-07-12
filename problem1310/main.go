package problem1310

import (
	"sort"
)

type ListElement struct {
	value int
	next  *ListElement
}

// Solution : assumes the input list doesn't have duplicated values and is not already sorted
// Complexity: O(n * log n) because of the sorting part
func Solution(in *ListElement) {
	sortedElements := make([]int, 0)

	e := in
	for e != nil {
		sortedElements = append(sortedElements, e.value)
		e = e.next
	}

	sort.Ints(sortedElements)

	i := 0
	low := true
	for e = in; e != nil; {
		if !low && i != (len(sortedElements)-1) {
			l := sortedElements[i]
			sortedElements[i] = sortedElements[i+1]
			sortedElements[i+1] = l
		}

		e.value = sortedElements[i]

		e = e.next
		i++
		low = !low
	}
}
