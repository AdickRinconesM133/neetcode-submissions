class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        const res = new Map()

        for (const w of strs) {
            const count = new Array(26).fill(0)

            for (const l of w) {
                count[l.charCodeAt(0) - "a".charCodeAt(0)]++
            }

            const key = count.join(",")

            if (!res.has(key)) {
                res.set(key, [])
            }

            res.get(key).push(w)
        }
        
        return Array.from(res.values())
    }
}
