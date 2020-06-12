package countriesInMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0, Solution([][]int{{}}))
	assert.Equal(t, 2, Solution([][]int{{1, 1, 1}, {2, 2, 2}}))
	assert.Equal(t, 3, Solution([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}))
	assert.Equal(t, 1, Solution([][]int{{1, 1, 1}, {1, 1, 1}}))
	assert.Equal(t, 1, Solution([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
	assert.Equal(t, 11, Solution([][]int{
		{5, 4, 4},
		{4, 3, 4},
		{3, 2, 4},
		{2, 2, 2},
		{3, 3, 4},
		{1, 4, 4},
		{4, 1, 1},
	}))
	assert.Equal(t, 1, Solution([][]int{
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
		{42, 42, 42},
	}))
	assert.Equal(t, 4, Solution([][]int{
		{1, 1, 1},
		{1, 2, 1},
		{2, 3, 3},
	}))
}
