func longestPalindromeSubseq(s string) int {
    dp := make([][]int, len(s))
    for i := range len(dp){
        dp[i] = make([]int, len(s))
    }
    return helper(s, 0, len(s)-1, dp)
}

func helper(s string, index1 int, index2 int, dp [][]int) int {
	if index1 > index2 || index2 >= len(s) || index1 >= len(s) {
		return 0
	}
	if index1 == index2 {
		return 1
	}
	if index2-index1 == 1 && s[index1] == s[index2] {
		return 2
	}
	if dp[index1][index2] != 0 {
		return dp[index1][index2]
	}
	if s[index1] == s[index2] {
		res := helper(s, index1+1, index2-1, dp)
		return res+2
    }
	x := helper(s, index1+1, index2, dp)
	y := helper(s, index1, index2-1, dp)
    dp[index1][index2] = max(x, y)
	return max(x, y)
}