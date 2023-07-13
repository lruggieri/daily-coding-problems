package problem1310

// Solution : O(n)
func Solution(randomF func(k int) int, deck [52]int) (shuffledDeck [52]int) {
	swap := func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	}

	for i := 0; i < len(deck); i++ {
		newPosition := randomF(len(deck))

		swap(i, newPosition)
	}

	return deck
}
