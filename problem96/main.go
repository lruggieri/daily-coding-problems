package problem96

func Solution(a []int) [][]int {
	solutions := make([][]int, 0)
	helper(a, 0, &solutions)
	return solutions
}

func helper(a []int, i int, solutions *[][]int) {
	if i == len(a)-1 {
		copyA := make([]int, len(a))
		copy(copyA, a)
		*solutions = append(*solutions, copyA)
		return
	}

	for j := i; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		helper(a, i+1, solutions)
		a[i], a[j] = a[j], a[i]
	}
}
