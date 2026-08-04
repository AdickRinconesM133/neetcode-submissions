func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	l := 0
	maxFreq := 0
	result := 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxFreq = max(maxFreq, count[s[r]])

		for r - l + 1 - maxFreq > k {
			count[s[l]]--
			l++
		}

		result = max(result, r - l + 1)
	}

	return result
}
