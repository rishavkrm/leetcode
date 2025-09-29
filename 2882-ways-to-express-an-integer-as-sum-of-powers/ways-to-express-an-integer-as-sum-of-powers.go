func numberOfWays(n int, x int) int {
	dp := make([][]int, n+1)
	for i := range len(dp) {
		dp[i] = make([]int, n+1)
	}
	return helper(n, x, 1, dp)%(1000000007)
}

func helper(n int, x int, curr int, dp [][]int) int {
	if n == 0 {
		return 1
	}
	if n < 0 || curr > n {
		return 0
	}
    if dp[n][curr] != 0{
        return dp[n][curr]
    }
	y1 := helper(n-pow(curr, x), x, curr+1, dp)
	y2 := helper(n, x, curr+1, dp)
    dp[n][curr] = y1+y2
	return y1 + y2
}

func pow(num int, pow int) int {
	res := 1
	for i := 0; i < pow; i++ {
		res *= num
	}
	return res
}