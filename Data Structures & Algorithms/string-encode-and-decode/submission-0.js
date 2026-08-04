class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        let encoded = ""

        for (const s of strs) {
            encoded += s.length+"#"+s
        }

        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        const res = new Array()
        let i = 0

        while (i < str.length) {
            let j = i
            while (str[j] != "#") {
                j++
            }
            const len = +str.slice(i, j)
            res.push(str.slice(j + 1, j + 1 + len))
            i = j + 1 + len
        }

        return res
    }
}
