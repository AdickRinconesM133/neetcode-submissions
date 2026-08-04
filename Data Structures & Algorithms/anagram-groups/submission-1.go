func groupAnagrams(strs []string) [][]string {
	anagramCounter := make(map[[26]int][]string, len(strs))

	for _, vstrs := range strs {
		alphKey := [26]int{}

		for _, vrune := range vstrs {
			alphKey[vrune - 'a']++
		}

		anagramCounter[alphKey] = append(anagramCounter[alphKey], vstrs) 
	}

	resultAnagrams := make([][]string, len(anagramCounter))
	anagramIndex := 0

	for _, v := range anagramCounter {
		resultAnagrams[anagramIndex] = v
		anagramIndex++
	}

	return resultAnagrams
}
