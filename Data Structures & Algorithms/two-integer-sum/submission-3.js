class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        const map = new Map()

        for (let i = 0; i < nums.length; i++) {
                const tar = target - nums[i]
                if (map.get(tar) != undefined) {
                    return [map.get(tar), i]
                } else {
                    map.set(nums[i], i)
                }
        }
    }
}
