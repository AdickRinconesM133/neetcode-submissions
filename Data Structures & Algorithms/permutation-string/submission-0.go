func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need := [26]int{}
	window := [26]int{}

	for l := 0; l < len(s1); l++ {
		need[s1[l]-'a']++
		window[s2[l]-'a']++
	}

	if need == window {
		return true
	}

	for r := len(s1); r < len(s2); r++ {
		window[s2[r]-'a']++
		window[s2[r-len(s1)]-'a']--

		if need == window {
		return true
		}
	}

	return false
}
