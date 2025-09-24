func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([][]int, len(nums))
	for i := range len(dp) {
		dp[i] = make([]int, target+1)
	}
	res := helper(nums, target, 0, &dp)
	if res < 0 {
		return -1
	}
	return res
}

func helper(nums []int, target int, index int, dp *[][]int) int {
	if target == 0 {
		return 0
	}
	if target < 0 || index >= len(nums) {
		return -9999
	}
	if (*dp)[index][target] != 0 {
		return (*dp)[index][target]
	}
	x := 1 + helper(nums, target-nums[index], index+1, dp)
	y := helper(nums, target, index+1, dp)
	(*dp)[index][target] = max(x, y)
	return max(x, y)
}