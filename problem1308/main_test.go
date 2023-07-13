package problem1310

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lruggieri/daily-coding-problems/util"
)

func TestSolution(t *testing.T) {
	t.Run("readme-example", func(t *testing.T) {
		pipes := map[string]map[string]int{
			"plant": {"A": 1, "B": 5, "C": 20},
			"A":     {"C": 15},
			"B":     {"C": 10},
			"C":     {},
		}

		assert.True(t, util.SliceSameUnordered(Solution(pipes), []Edge{
			{From: "plant", To: "A", Cost: 1},
			{From: "plant", To: "B", Cost: 5},
			{From: "B", To: "C", Cost: 10},
		}))
	})
}
