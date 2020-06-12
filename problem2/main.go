package main

//complexity O(n^2)
func Trivial(iSlice []int) []int {
	newSlice := make([]int, len(iSlice))
	for i := range iSlice {
		newSlice[i] = 1
	}

	for i := range iSlice {
		for j, k := range iSlice {
			if j != i {
				newSlice[i] *= k
			}
		}
	}

	return newSlice
}

//complexity O(n)
//it fails with edge cases with 0 values
func WithDivision(iSlice []int) []int {
	newSlice := make([]int, len(iSlice))

	product := 1
	for _, v := range iSlice {
		product *= v
	}

	for i, v := range iSlice {
		if v == 0 {
			newSlice[i] = 0
		} else {
			newSlice[i] = product / v
		}
	}

	return newSlice
}

//complexity O(n)
func WithoutDivision(iSlice []int) []int {
	newSlice := make([]int, len(iSlice))

	type rightLeftValue struct {
		leftMultiplication  int //contains multiplications of all elements to the left without the current index
		rightMultiplication int //contains multiplications of all elements to the right without the current index
	}
	newTempSlice := make([]rightLeftValue, len(iSlice))

	for iLeft, iRight := 0, len(iSlice)-1; iLeft < len(iSlice) && iRight >= 0; iLeft, iRight = iLeft+1, iRight-1 {
		if iLeft == 0 {
			newTempSlice[iLeft].leftMultiplication = 1
		} else {
			newTempSlice[iLeft].leftMultiplication = newTempSlice[iLeft-1].leftMultiplication * iSlice[iLeft-1]
		}
		if iRight == len(iSlice)-1 {
			newTempSlice[iRight].rightMultiplication = 1
		} else {
			newTempSlice[iRight].rightMultiplication = newTempSlice[iRight+1].rightMultiplication * iSlice[iRight+1]
		}
	}

	for i := range iSlice {
		newSlice[i] = newTempSlice[i].rightMultiplication * newTempSlice[i].leftMultiplication
	}

	return newSlice
}
