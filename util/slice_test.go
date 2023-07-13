package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSameUnordered(t *testing.T) {
	assert.True(t, SliceSameUnordered([]int{}, []int{}))
	assert.True(t, SliceSameUnordered([]int{1, 2, 3}, []int{1, 3, 2}))
	assert.False(t, SliceSameUnordered([]int{3, 3, 3}, []int{1, 3, 2}))
	assert.False(t, SliceSameUnordered([]int{1}, []int{1, 2}))
}
