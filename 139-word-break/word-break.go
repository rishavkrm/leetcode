func wordBreak(s string, wordDict []string) bool {
	return helper(s, wordDict, "", map[string]int{})
}

func helper(s string, w []string, curr string, dp map[string]int) bool {
	if s == curr {
		return true
	}
	if len(curr) >= len(s) {
		return false
	}
	if s[:len(curr)] != curr {
		return false
	}
    if dp[curr] == -1{
        return false
    }
	for i := 0; i < len(w); i++ {
		res1 := helper(s, w, curr+w[i], dp)
		if res1 {
			return true
		}

	}
	dp[curr] = -1
	return false
}