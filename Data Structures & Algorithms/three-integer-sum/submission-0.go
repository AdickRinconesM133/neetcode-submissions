func threeSum(nums []int) [][]int {
	sol := [][]int{}

	sort.Ints(nums)

	for i, a := range nums {
		if i > 0 && a == nums[i - 1] {
			continue
		}

		left, right := i + 1, len(nums) - 1

		for left < right {
			sum := a + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				sol = append(sol, []int{a, nums[left], nums[right]})
				left++
				for nums[left] == nums[left - 1] && left < right {
					left++
				}
			}
		}
	}

	return sol
}
