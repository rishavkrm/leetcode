func longestPalindrome(s string) string {
	dp := make([][]int, len(s))
	for i := range len(dp) {
		dp[i] = make([]int, len(s))
	}
	size := helper(s, 0, len(s)-1, dp)
	fmt.Println(size)
	for i := 0; i <= len(s)-size; i++ {
		curr := s[i : i+size]
		if reverse(curr) == curr {
			return curr
		}
	}
	return ""
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
		if res == index2-index1-1 {
            dp[index1][index2] = res+2
			return res + 2
		}
    }
	x := helper(s, index1+1, index2, dp)
	y := helper(s, index1, index2-1, dp)
    dp[index1][index2] = max(x, y)
	return max(x, y)
}

func reverse(s string) string {
	temp := ""
	for i := 0; i < len(s); i++ {
		temp += string(s[len(s)-i-1])
	}
	return temp
}

// aaaaaaaaaaaaaaaaaaaaaaaaa