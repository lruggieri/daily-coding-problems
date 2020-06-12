package problem241

import "sort"

// Reverse sort the input, then loop through it and when the current index is not at least the current value, we already know that we do not need to go on
// O(N*logN), due to sorting
func Solution(papersResults []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(papersResults))) // sorted in reverse

	if len(papersResults) == 0 || papersResults[0] == 0 {
		return 0
	}

	hIndex := 0
	for i, n := range papersResults {
		if n < i+1 {
			break
		}
		hIndex++
	}
	return hIndex
}
