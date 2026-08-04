func productExceptSelf(nums []int) []int {
 pre := make([]int, len(nums))
 post := make([]int, len(nums))
 output := make([]int, len(nums))
 prod := 1

  for i := 0; i <= len(nums) - 1; i++ {
	prod *= nums[i]
	pre[i] = prod
 }

prod = 1

  for i := len(nums) - 1; i >= 0; i-- {
	prod *= nums[i]
	post[i] = prod
 }

  for i := 0; i <= len(nums) - 1; i++ {
    izq := 1
    if i > 0 {
        izq = pre[i-1]
    }

    der := 1
    if i < len(nums)-1 {
        der = post[i+1]
    }

    output[i] = izq * der
 }

 return output
}
