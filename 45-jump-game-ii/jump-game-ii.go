func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = -2
	}
    
	x := helper(nums, 0, 0, dp)
    fmt.Println(dp)
    return x    
}

func helper(nums []int, index int, curr int, dp []int) int {
	if index >= len(nums) {
		return 1000000
	}
	if index == len(nums)-1 {
		return curr
	}
	if dp[index] != -2 {
		return dp[index] + curr
	}
	x := nums[index] //2
	mini := len(nums) // 2
	for i := 1; i <= x; i++ {
		mini = min(mini, helper(nums, index+i, curr+1, dp))
	}
	dp[index] = mini - curr
	return mini
}