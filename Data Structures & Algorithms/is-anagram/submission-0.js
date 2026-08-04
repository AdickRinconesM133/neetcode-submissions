class Solution {
    isAnagram(s, t) {
        const fs = new Map();
        const ft = new Map();
 
        if (s.length != t.length) {
            return false
        }

        for (let i = 0; i < s.length; i++) {
            fs.set(s[i], 1 + (fs.get(s[i]) ?? 0))
            ft.set(t[i], 1 + (ft.get(t[i]) ?? 0))
        }

        for (const l of s) {
            if (fs.get(l) != ft.get(l)) {
                return false 
            }
        }

        return true
    }
}
