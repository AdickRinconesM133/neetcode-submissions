func isPalindrome(s string) bool {
    left, right := 0, len(s) - 1 

    for left < right {
        rLeft := rune(s[left])
        rRight := rune(s[right])

        if !unicode.IsLetter(rLeft) && !unicode.IsDigit(rLeft) {
            left++
            continue
        }

        if !unicode.IsLetter(rRight) && !unicode.IsDigit(rRight) {
            right--
            continue
        }

        if unicode.ToLower(rLeft) != unicode.ToLower(rRight) {
            return false
        } else {
            left++
            right--
        }
    }

    return true
}
