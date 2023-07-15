package problem1310

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lruggieri/daily-coding-problems/util"
)

func TestSolution(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		config := map[string][]string{
			"A": {"B"},
			"B": {"C"},
			"C": {"A", "D"},
			"D": {"E", "F"},
			"E": {"C", "F", "I", "J"},
			"F": {"G"},
			"G": {},
			"H": {},
			"I": {"H"},
			"J": {},
		}

		g := buildGraphFromConfig(config)

		//g.GetNode("A").Print()

		assert.Equal(t, 10, g.GetNode("A").ConnectedNodes())
		assert.Equal(t, 1, g.GetNode("A").ConnectedNodes("AB"))
		assert.Equal(t, 2, g.GetNode("A").ConnectedNodes("BC"))
		assert.Equal(t, 10, g.GetNode("D").ConnectedNodes())
		assert.Equal(t, 1, g.GetNode("G").ConnectedNodes())

		assert.True(t, util.SliceSameUnordered([]string{"EI", "EJ", "FG", "IH"}, Solution(g)))

	})
}
