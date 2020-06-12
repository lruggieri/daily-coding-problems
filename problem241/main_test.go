package problem241

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, Solution([]int{4, 3, 0, 1, 5}))
	assert.Equal(t, 1, Solution([]int{0, 0, 0, 1}))
	assert.Equal(t, 0, Solution([]int{0, 0, 0, 0}))
	assert.Equal(t, 5, Solution([]int{5, 6, 7, 8, 9}))
	//assert.Equal(t, 5, Solution([]int{1, 10, 9, 8, 345, 6, 87, 0, 34}))
}
