func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, v := range nums {
        _, exist := seen[v]

        if exist {
            return true
        }
        
        seen[v] = true
    }

    return false
}
