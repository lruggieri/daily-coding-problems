package problem893

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("no-color-overlap", func(t *testing.T) {
		assert.Equal(t, 4, Solution([][]int{
			{1, 2, 3, 4},
			{2, 3, 4, 1},
			{3, 4, 1, 2},
			{4, 1, 2, 3},
		}))
	})
	t.Run("color-overlap", func(t *testing.T) {
		assert.Equal(t, 6, Solution([][]int{
			{2, 1, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 4, 3},
			{4, 1, 2, 3},
		}))
	})
	t.Run("same-as-previous-with-last-overlap", func(t *testing.T) {
		assert.Equal(t, 5, Solution([][]int{
			{2, 1, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 4, 3},
			{1, 4, 2, 3},
		}))
	})
}
