package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblem2(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, Trivial([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{24, 12, 8, 6}, WithDivision([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{24, 12, 8, 6}, WithoutDivision([]int{1, 2, 3, 4}))

	assert.Equal(t, []int{24, 8, 12, 6}, Trivial([]int{1, 3, 2, 4}))
	assert.Equal(t, []int{24, 8, 12, 6}, WithDivision([]int{1, 3, 2, 4}))
	assert.Equal(t, []int{24, 8, 12, 6}, WithoutDivision([]int{1, 3, 2, 4}))
}
