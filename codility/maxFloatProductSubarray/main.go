package maxFloatProductSubarray

import (
	"math"
)

func Solution(A []float64) float64 {

	//max positive value, min negative value, maximum float between zero and one
	maxPositive, minNegative, zeroOneMaxFloat := 1.0, 1.0, -10000.0

	globalMax := 1.0       // our final result
	nonFloatFound := false //if we have all float values, we still have to return the greatest among them

	for _, n := range A {
		if n > 1 {
			maxPositive = maxPositive * n
			minNegative = math.Min(minNegative*n, 1)
		} else if n >= 0 && n < 1 {
			// this number can only worsen the actual result!
			maxPositive = 1.0
			minNegative = 1.0
			zeroOneMaxFloat = math.Max(zeroOneMaxFloat, n)
		} else {
			// n < 0
			/*
				at this point we have
				minNegative <= 1 <= maxPositive
			*/
			oldPositive := maxPositive
			maxPositive = math.Max(minNegative*n, 1)
			minNegative = oldPositive * n
		}

		if maxPositive > globalMax {
			nonFloatFound = true
			globalMax = maxPositive
		}
	}
	if !nonFloatFound {
		globalMax = zeroOneMaxFloat
	}

	return globalMax
}
