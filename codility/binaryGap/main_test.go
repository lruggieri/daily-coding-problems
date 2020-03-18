package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBinaryRepresentation(t *testing.T) {
	assert.Equal(t, "101", GetBinaryRepresentation(5,2))
	assert.Equal(t, "12", GetBinaryRepresentation(5,3))
	assert.Equal(t, "11", GetBinaryRepresentation(5,4))
	assert.Equal(t, "10", GetBinaryRepresentation(5,5))
	assert.Equal(t, "5", GetBinaryRepresentation(5,11))
	assert.Equal(t, "1011", GetBinaryRepresentation(11,2))
	assert.Equal(t, "0", GetBinaryRepresentation(0,2))
	assert.Equal(t, "1000010", GetBinaryRepresentation(66,2))
}

func TestSolution(t *testing.T) {
	assert.Equal(t, 0,Solution(0))
	assert.Equal(t, 0,Solution(2))
	assert.Equal(t, 1,Solution(11))
	assert.Equal(t, 2,Solution(9))
	assert.Equal(t, 4,Solution(66))
}