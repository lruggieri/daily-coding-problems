package maxFloatProductSubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2.0, Solution([]float64{0.7, -1, -2, 0.5}))
	assert.Equal(t, 0.2, Solution([]float64{0.1, 0.2}))
	assert.Equal(t, 0.2, Solution([]float64{0.1, 0.2}))
	assert.Equal(t, 0.2, Solution([]float64{-10, 0.1, 0.2}))
	assert.Equal(t, 0.0, Solution([]float64{0, 0, 0, 0}))
	assert.Equal(t, 65.6, Solution([]float64{-7, 2, 4.1, 8}))
}
