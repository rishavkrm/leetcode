func wordBreak(s string, wordDict []string) []string {
    result := []string{}
    helper(s, wordDict, "", map[string]int{}, []string{}, &result)
    return result
}
func helper(s string, w []string, curr string, dp map[string]int, arr []string, result *[]string) bool{
    // fmt.Println(curr)
	if s == curr {
        *result = append(*result, arrToStr(arr))
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
    res := false
	for i := 0; i < len(w); i++ {
        arr = append(arr, w[i])
		res = helper(s, w, curr+w[i], dp, arr, result) || res
        arr = arr[:len(arr)-1]

	}
    if res{
        return res
    }
	dp[curr] = -1
    return res
}

func arrToStr(strs []string)string{
    res := ""
    for _,str := range strs{
        res += " "
        res += str
    }

    if len(res)>0{
        res = res[1:]
    }
    return res
}