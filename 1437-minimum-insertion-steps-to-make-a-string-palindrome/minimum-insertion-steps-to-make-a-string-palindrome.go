func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := range len(dp) {
		dp[i] = make([]int, len(dp))
	}
    for i := range len(dp){
        for j := range len(dp[0]){
            dp[i][j] = -1
        }
    }
	return helper(s, 0, len(s)-1, dp)
}

func helper(s string, left int, right int, dp [][]int) int {
	if left >= right {
		return 0
	}
	if dp[left][right] != -1 {
		return dp[left][right]
	}
	if s[left] == s[right] {
		return helper(s, left+1, right-1, dp)
	}
	x := 1 + helper(s, left+1, right, dp)
	y := 1 + helper(s, left, right-1, dp)
    dp[left][right] = min(x, y)
	return min(x, y)
}