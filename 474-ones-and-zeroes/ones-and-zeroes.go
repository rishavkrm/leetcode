func findMaxForm(strs []string, m int, n int) int {
	counts := make([][]int, len(strs))
	for i := range len(counts) {
		counts[i] = make([]int, 2)
	}
	for i, str := range strs {
		for j := range len(str) {
			if str[j] == '0' {
				counts[i][0] += 1
			} else if str[j] == '1' {
				counts[i][1] += 1
			}
		}
	}

	dpDict := map[int][][]int{}
	for i := range len(strs) {
		dp := make([][]int, m+1)
		for i := range len(dp) {
			dp[i] = make([]int, n+1)
		}
		for i := range len(dp) {
			for j := range len(dp[0]) {
				dp[i][j] = -1
			}
		}
		dpDict[i] = dp
	}
	return helper(counts, 0, m, n, dpDict)
}

func helper(count [][]int, index int, m int, n int, dp map[int][][]int) int {
	if index >= len(count) || m < 0 || n < 0 {
		return 0
	}
	if dp[index][m][n] != -1 {
		return dp[index][m][n]
	}
	zero, one := count[index][0], count[index][1]
	x, y := 0, 0
	if m-zero >= 0 && n-one >= 0 {
		x = 1 + helper(count, index+1, m-zero, n-one, dp)
	}
	y = helper(count, index+1, m, n, dp)
	dp[index][m][n] = max(x, y)
	return max(x, y)
}
