class Solution {
    hasDuplicate(nums) {
        const numbers = new Set();

        for (let n of nums) {
            if (numbers.has(n)) {
                return true;
            } else {
                numbers.add(n)
            }
        }

        return false;
    }
}