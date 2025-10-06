func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
    dp := make([]int, len(nums))
    for i := range len(dp){
        dp[i] = -1
    }
	dict := map[int]int{}
    freq := map[int]int{}
	for i := range len(nums) {
		dict[nums[i]] = i
        freq[nums[i]] += 1
	}
	return helper(nums, 0, dict, freq, dp)
}

func helper(nums []int, index int, dict map[int]int, freq map[int]int, dp []int) int {
	if index >= len(nums) {
		return 0
	}
    if dp[index] != -1{
        return dp[index]
    } 
	// case 1 - take
	nextIndex := dict[nums[index]]+1
	i, exists := dict[nums[index]+1]
	if exists {
		nextIndex = i + 1
	}
	x := nums[index]*freq[nums[index]] + helper(nums, nextIndex, dict, freq, dp)
	// case 2 - not take
	y := helper(nums, index+1, dict, freq, dp)
    dp[index] = max(x, y)
	return max(x, y)
}