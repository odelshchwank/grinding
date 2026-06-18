func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)
	rs := []rune(s)
	rt := []rune(t)

	for _, r := range rs {
		counts[r]++
	}

	for _, r := range rt {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}

	return true

}
