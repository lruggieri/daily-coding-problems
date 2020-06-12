package problem158

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, Solution(
		[][]bool{
			{true, true, false},
			{true, true, false},
			{false, true, true},
		},
	))
	assert.Equal(t, 10, Solution(
		[][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
			{true, true, true},
		},
	))
	assert.Equal(t, 7, Solution(
		[][]bool{
			{true, true, true, true},
			{true, true, true, false},
			{true, false, true, true},
			{true, true, true, true},
		},
	))
	assert.Equal(t, 1, Solution(
		[][]bool{
			{true, true, true, true},
			{false, false, false, true},
			{false, false, false, true},
			{false, false, false, true},
		},
	))
	assert.Equal(t, 20, Solution(
		[][]bool{
			{true, true, true, true},
			{true, true, true, true},
			{true, true, true, true},
			{true, true, true, true},
		},
	))
	// if we cannot reach the end (false at bottom right)
	assert.Equal(t, 0, Solution(
		[][]bool{
			{true, true, true, true},
			{true, true, true, false},
			{true, false, true, true},
			{true, true, true, false},
		},
	))
}
