func hasDuplicate(nums []int) bool {
    numExist := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := numExist[v]; ok {
			return true
		}

		numExist[v] = struct{}{}
	}

	return false
}
