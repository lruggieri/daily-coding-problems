package problem893

import "math"

/*
	{2, 1, 3, 4}
	{1, 2, 3, 4},
	{1, 2, 4, 3},
	{4, 1, 2, 3},
*/

// Solution :
func Solution(in [][]int) int {

	previousMin, previousSecondMin, previousColor := 0, 0, -1

	for house, houseColors := range in {
		currentMin, currentSecondMin, currentColor := math.MaxInt32, math.MaxInt32, -1
		for color, _ := range houseColors {
			cost := in[house][color]
			if color == previousColor {
				cost += previousSecondMin
			} else {
				cost += previousMin
			}

			if cost < currentMin {
				currentSecondMin = currentMin
				currentMin = cost
				currentColor = color
			} else if cost < currentSecondMin {
				currentSecondMin = cost
			}
		}

		previousMin, previousSecondMin, previousColor = currentMin, currentSecondMin, currentColor
	}

	return previousMin
}
