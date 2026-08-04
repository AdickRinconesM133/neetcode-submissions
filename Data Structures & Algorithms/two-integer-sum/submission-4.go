func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, v := range nums {
        rest := target - v

        if is, exist := seen[rest]; exist {
            return []int{is, i}
        }

        seen[v] = i
    }

    return nil
}
