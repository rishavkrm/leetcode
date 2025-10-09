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

func helper(word1 string, word2 string, index1 int, index2 int, dp [][]int) int {
	if index1 >= len(word1) && index2 >= len(word2) {
		return 0
	}
	if index1 >= len(word1) {
		return len(word2) - index2
	}
	if index2 >= len(word2) {
		return len(word1) - index1
	}
	if dp[index1][index2] != -1 {
		return dp[index1][index2]
	}
	if word1[index1] == word2[index2] {
		dp[index1][index2] = helper(word1, word2, index1+1, index2+1, dp)
		return dp[index1][index2]
	}
	x, y := 501, 501
	if index1 < len(word1) {
		x = 1 + helper(word1, word2, index1+1, index2, dp)
	}
	if index2 < len(word2) {
		y = 1 + helper(word1, word2, index1, index2+1, dp)
	}
	dp[index1][index2] = min(x, y)
	return min(x, y)

}