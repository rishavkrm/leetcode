func numSquares(n int) int {
	dp := make([][]int, n+1)
	for i := range len(dp) {
		dp[i] = make([]int, 101)
	}

	return helper(n, 1, dp)
}

func helper(n int, curr int, dp [][]int) int {
	if n == 0 {
		return 0
	}
	if curr > 100 || n < 0 {
		return 9999
	}
	if dp[n][curr] != 0 {
		return dp[n][curr]
	}
	x1 := 1 + helper(n-curr*curr, curr, dp)
	x2 := helper(n, curr+1, dp)
	dp[n][curr] = min(x1, x2)
	return min(x1, x2)
}
