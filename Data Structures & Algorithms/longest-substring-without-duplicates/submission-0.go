func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]struct{})
	longestCount := 0

	for l, r := 0, 0; r < len(s); r++ {
		for _, exist := charSet[s[r]]; exist; _, exist = charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}

		charSet[s[r]] = struct{}{}
		longestCount = max(longestCount, r-l+1) 
	} 

	return longestCount
}
