class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const chatFreq = new Map()
        const freqArray = new Array(nums.length + 1)
        const resArray = new Array()

        for (const n of nums) {
            chatFreq.set(n, 1 + (chatFreq.get(n) ?? 0))
        }

        for (const [key, v] of chatFreq) {
            if (freqArray[v] == undefined) {
                freqArray[v] = []
            }
            freqArray[v].push(key)
        }
        
        for (let i = freqArray.length - 1; i >= 0; i--) {
            if (freqArray[i] != undefined) {
                for (const val of freqArray[i]) {
                    resArray.push(val)
                    if (resArray.length == k) {
                        return resArray
                    }
                }
            }
        }
        return resArray
    }
}
