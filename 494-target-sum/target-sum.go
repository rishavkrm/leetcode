func findTargetSumWays(nums []int, target int) int {
	dp := make([][]int, len(nums))
	for i := range len(dp) {
		dp[i] = make([]int, 5000)
	}
	return helper(nums, target, 0, &dp)
}

func helper(nums []int, target int, index int, dp *[][]int) int {
	if index == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	if index >= len(nums) {
		return 0
	}
	if (*dp)[index][target+2000] != 0 {
		return (*dp)[index][target+2000]
	}
	x := helper(nums, target-nums[index], index+1, dp)
	y := helper(nums, target+nums[index], index+1, dp)
	(*dp)[index][target+2000] = x + y
	return x + y

}