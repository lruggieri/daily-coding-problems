package problem1052

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lruggieri/daily-coding-problems/util"
)

func TestSolution(t *testing.T) {
	t.Run("solvable", func(t *testing.T) {
		teamA, teamB, solvable := Solution(map[int][]int{
			0: {3},
			1: {2},
			2: {1, 4},
			3: {0, 4, 5},
			4: {2, 3},
			5: {3},
		})

		require.True(t, solvable)
		if len(teamA) == 4 {
			assert.True(t, util.SliceSameUnordered(teamA, []int{0, 1, 4, 5}))
			assert.True(t, util.SliceSameUnordered(teamB, []int{2, 3}))
		} else {
			assert.True(t, util.SliceSameUnordered(teamB, []int{0, 1, 4, 5}))
			assert.True(t, util.SliceSameUnordered(teamA, []int{2, 3}))
		}
	})

	t.Run("unsolvable", func(t *testing.T) {
		teamA, teamB, solvable := Solution(map[int][]int{
			0: {3},
			1: {2},
			2: {1, 3, 4},
			3: {0, 2, 4, 5},
			4: {2, 3},
			5: {3},
		})

		require.False(t, solvable)
		assert.Empty(t, teamA)
		assert.Empty(t, teamB)
	})
}
