package graph_loops

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lruggieri/daily-coding-problems/util"
)

func TestSolution(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		loops := Solution(map[int][]int{
			1: {3},
			2: {1, 3},
			3: {2, 4, 8},
			4: {5},
			5: {3, 4},
			8: {9},
			9: {4},
		})

		assert.True(t, util.MatrixSameUnordered(loops, [][]int{
			{1, 3, 2},
			{3, 2},
			{3, 4, 5},
			{4, 5},
			{3, 8, 9, 4, 5},
		}))
	})

	t.Run("chain", func(t *testing.T) {
		loops := Solution(map[int][]int{
			1: {2},
			2: {1, 3},
			3: {2, 4},
			4: {3, 5},
			5: {4, 6},
			6: {5},
		})

		assert.True(t, util.MatrixSameUnordered(loops, [][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{4, 5},
			{5, 6},
		}))
	})

	t.Run("broken-chain", func(t *testing.T) {
		loops := Solution(map[int][]int{
			1: {2},
			2: {1, 3},
			3: {4}, // doesn't connect back to 2
			4: {3, 5},
			5: {4, 6},
			6: {5},
		})

		assert.True(t, util.MatrixSameUnordered(loops, [][]int{
			{1, 2},
			{3, 4},
			{4, 5},
			{5, 6},
		}))
	})
}
