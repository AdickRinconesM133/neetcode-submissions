func isAnagram(s string, t string) bool {
    charCountS := make(map[rune]int)
    charCountT := make(map[rune]int)

    if len(s) != len(t) {
        return false
    }

    for _, v := range s {
        charCountS[v] += 1
    }

    for _, v := range t {
        charCountT[v] += 1
    }  

    for i, v := range charCountS {
        if charCountT[i] != v {
            return false
        }
    }

    return true
}
