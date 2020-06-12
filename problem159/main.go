package problem159

func Solution(s string) rune {
	alreadySeen := make(map[rune]bool)
	for _, c := range s {
		if alreadySeen[c] {
			return c
		}
		alreadySeen[c] = true
	}
	return -1
}
