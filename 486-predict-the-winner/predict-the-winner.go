func predictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := range len(dp) {
		dp[i] = make([]int, len(nums))
	}
	for i := range len(dp) {
		for j := range len(dp[i]) {
			dp[i][j] = -1
		}
	}
	xx := helper(nums, 0, len(nums)-1, dp)
	return xx >= 0
}

func helper(nums []int, left int, right int, dp [][]int) int {
	if left > right {
		return 0
	}
	if dp[left][right] != -1{
	    return dp[left][right]
	}
	x := nums[left] - helper(nums, left+1, right, dp)
    y := nums[right] - helper(nums, left, right-1, dp)
    dp[left][right] = max(x, y)
	return dp[left][right]
}