func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1))
	for i := range len(dp) {
		dp[i] = make([]int, len(word2))
	}
    for i := range len(dp){
        for j := range len(dp[0]){
            dp[i][j] = -1
        }
    }
	return helper(word1, word2, 0, 0, dp)
}

func helper(word1 string, word2 string, i1 int, i2 int, dp [][]int) int {
	if i1 >= len(word1) || i2 >= len(word2) {
		if i1 >= len(word1) {
			return len(word2) - i2 
		}
		return len(word1) - i1
	}
    if dp[i1][i2] != -1{
        return dp[i1][i2]
    }
	if word1[i1] == word2[i2] {
		return helper(word1, word2, i1+1, i2+1, dp)
	}
	// insert or delete
	x := 1 + helper(word1, word2, i1+1, i2, dp)
	y := 1 + helper(word1, word2, i1, i2+1, dp)
	// replace
	z := 1 + helper(word1, word2, i1+1, i2+1, dp)
    dp[i1][i2] = min(min(x, y), z)
	return min(min(x, y), z)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}