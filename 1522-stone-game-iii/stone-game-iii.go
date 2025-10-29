func stoneGameIII(stoneValue []int) string {
	dp := make([]int, len(stoneValue))
	for i := range len(dp) {
		dp[i] = -10000
	}
	x := helper(stoneValue, 0, dp)
	if x > 0 {
		return "Alice"
	} else if x < 0 {
		return "Bob"
	} else {
		return "Tie"
	}

}

func helper(arr []int, index int, dp []int) int {
	if index >= len(arr) {
		return 0
	}
	if dp[index] != -10000 {
		return dp[index]
	}
	res := []int{-10000, -10000, -10000}
	curr := 0
	for i := range 3 {
		if i+index >= len(arr) {
			continue
		}
		curr += arr[i+index]
		x := curr - helper(arr, index+i+1, dp)
		res[i] = x
	}
	dp[index] = max(res[0], max(res[1], res[2]))
	return max(res[0], max(res[1], res[2]))
}