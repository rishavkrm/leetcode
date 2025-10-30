func winnerSquareGame(n int) bool {
	dp := make([]int, n+1)
	for i := range len(dp) {
		dp[i] = -1
	}
	xx := helper(n, dp)
    return xx % 2 == 1
}

func helper(n int, dp []int) int {
	if n <= 0 {
		return 0
	}
	if dp[n] != -1 {
		return dp[n]
	}
    xx := 0
	for i := 1; i <= 317; i++ {
		if i*i > n {
			break
		}
		xx = 1 + helper(n-i*i, dp)
        if xx % 2 == 1{
            dp[n] = xx
            break
        }
	}
	dp[n] = xx
	return xx
}