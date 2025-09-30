func minSteps(n int) int {
    dp := make([][]int, n)
    for i := range len(dp){
        dp[i] = make([]int, n)
    }
	return helper(n, 1, 0, dp)
}

func helper(n int, curr int, copied int, dp [][]int) int {
	if curr == n {
		return 0
	}
	if curr > n {
		return 10000
	}
    if dp[curr][copied] != 0{
        return dp[curr][copied]
    }
	x1 := 10000
	x2 := 10000
	if copied == 0 {
		x1 = 1 + helper(n, curr, curr, dp)
	} else if curr == copied {
		x2 = 1 + helper(n, curr+copied, copied, dp)
	} else {
		x1 = 1 + helper(n, curr, curr, dp)
		x2 = 1 + helper(n, curr+copied, copied, dp)
	}
    dp[curr][copied] = min(x1, x2)
	return min(x1, x2)
}