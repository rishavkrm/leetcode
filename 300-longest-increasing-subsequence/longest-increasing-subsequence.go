func lengthOfLIS(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range len(dp) {
		dp[i] = make([]int, len(nums)+1)
	}
	return helper(nums, 0, -1, dp)
}

func helper(nums []int, index int, curr int, dp [][]int) int {
	if index >= len(nums) {
		return 0
	}
	if dp[index][curr+1] != 0 {
		return dp[index][curr+1]
	}
	//  not take
	x, y := 0, 0
	x = helper(nums, index+1, curr, dp)
	// if bigger take
	if curr < 0 || nums[index] > nums[curr] {
		y = 1 + helper(nums, index+1, index, dp)
	}
	dp[index][curr+1] = max(x, y)
	return max(x, y)
}