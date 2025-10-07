func maxSubArray(nums []int) int {
	prefix := make([]int, len(nums))
	curr := 0
	for i := 0; i < len(nums); i++ {
		prefix[i] = curr + nums[i]
		curr = prefix[i]
	}
	mini := 0
	maxi := prefix[0]
	for i := range len(prefix) {
		maxi = max(maxi, prefix[i]-mini)
		mini = min(mini, prefix[i])

	}
	return maxi

}

// [-2, -1, -4, 0, -1, 1, 2, -3, 1]
// [-2, -3]