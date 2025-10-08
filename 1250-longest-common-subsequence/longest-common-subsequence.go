func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1))
    for i := range len(dp){
        dp[i] = make([]int, len(text2))
    }
    for i := 0; i < len(dp); i++{
        for j := 0; j < len(dp[0]); j++{
            dp[i][j] = -1
        }
    }
    return helper(text1, text2, 0, 0, dp)
}

func helper(text1 string, text2 string, index1 int, index2 int, dp[][]int) int {
	if index1 >= len(text1) || index2 >= len(text2) {
		return 0
	}
    if dp[index1][index2] != -1{
        return dp[index1][index2]
    }
	x, y, z := 0, 0, 0
	if text1[index1] == text2[index2] {
		x = 1 + helper(text1, text2, index1+1, index2+1, dp)
	} else {
		y = helper(text1, text2, index1+1, index2, dp)
		z = helper(text1, text2, index1, index2+1, dp)
	}
    dp[index1][index2] = max(x, max(y, z))
	return max(x, max(y, z))
}   