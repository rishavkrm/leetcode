func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    for i := range(len(dp)){
        dp[i] = -1
    }
	return helper(nums, target, dp)
}

func helper(nums []int, target int, dp []int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
    if dp[target] != -1{
        return dp[target]
    }
	res := 0
	for index := range len(nums) {
		res += helper(nums, target-nums[index],dp)
	}
    dp[target] = res
	return res

}