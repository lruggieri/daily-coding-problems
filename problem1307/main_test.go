package problem1310

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		randomFunc := func(k int) int {
			return rand.Intn(k)
		}

		deck := [52]int{}
		for i := 0; i < len(deck); i++ {
			deck[i] = i + 1
		}

		shuffledDeck := Solution(randomFunc, deck)

		assert.NotEqual(t, deck, shuffledDeck)
	})
}
