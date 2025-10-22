func numDistinct(s string, t string) int {
    dp := make([][]int, len(s))
	for i := range len(dp) {
		dp[i] = make([]int, len(t))
	}
    for i := 0; i < len(dp); i++{
        for j := 0; j < len(dp[0]); j++{
            dp[i][j] = -1
        }
    }
    return helper(s, t, 0, 0, dp)
}

func helper(s string, t string, i1 int, i2 int, dp [][]int) int {
	if i2 == len(t) {
		return 1
	}
    if i1 >= len(s) || i2 >= len(t){
        return 0
    }
    if dp[i1][i2] != -1{
        return dp[i1][i2]
    }
	x, y := 0, 0
	if s[i1] == t[i2] {
		x = helper(s, t, i1+1, i2, dp)
		y = helper(s, t, i1+1, i2+1, dp)
	} else {
		x = helper(s, t, i1+1, i2, dp)
	}
    dp[i1][i2] = x + y
    return x + y
}