package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToK(t *testing.T) {
	assert.Equal(t, true, AddToK([]int{5,1,9,7,8,6},13))
	assert.Equal(t, true, AddToK([]int{5,1,9,7,8,6},11))
	assert.Equal(t, true, AddToK([]int{5,1,9,7,8,599},600))
	assert.Equal(t, true, AddToK([]int{5,7,8,9,8},16))
	assert.Equal(t, true, AddToK([]int{51,2,3,4,5,6},8))
}
