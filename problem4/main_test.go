package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLowestPositiveMissingInteger(t *testing.T) {

	assert.Equal(t,2,FindLowestPositiveMissingInteger([]int{3,4,-1,1}))
	assert.Equal(t,5,FindLowestPositiveMissingInteger([]int{3,4,-1,2,1}))
	assert.Equal(t,1,FindLowestPositiveMissingInteger([]int{3,4,-1,3}))
	assert.Equal(t,1,FindLowestPositiveMissingInteger([]int{-5,-3,-10,-1})) // all negative
	assert.Equal(t,1,FindLowestPositiveMissingInteger([]int{-5,-3,7,9})) // negative and positive
	assert.Equal(t,5,FindLowestPositiveMissingInteger([]int{1,2,4,3}))
	assert.Equal(t,6,FindLowestPositiveMissingInteger([]int{-1,1,1,2,3,4,5,2,1,1,2,0,0})) // lots of repetition

}
