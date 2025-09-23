func canPartition(nums []int) bool {
    total := 0
	for i := range nums {
		total += nums[i]
	}
	if total%2 == 1 {
		return false
	}
	dp := make([][]int, len(nums))
	for i := range len(dp) {
		dp[i] = make([]int, total/2)
	}
	
	return helper(nums, total/2, 0, 0, dp)
}

func helper(nums []int, target int, index int, curr int, dp [][]int) bool {
	if curr == target {
		return true
	}
	if index >= len(nums) || curr > target {
		return false
	}
	if dp[index][curr] != 0 {
		return dp[index][curr] > 0
	}
	x := helper(nums, target, index+1, curr, dp)
	y := helper(nums, target, index+1, curr+nums[index], dp)
	if x || y {
		dp[index][curr] = 1
	} else {
		dp[index][curr] = -1
	}
	return x || y
}