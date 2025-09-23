func partition(s string) [][]string {
    res := [][]string{}
    helper(s, 0, []string{}, 0, &res)
    return res
}

func helper(s string, index int, temp []string, used int, res *[][]string) {
    if len(s) == used {
		(*res) = append(*res, copyArray(temp))
		return
	}
	if index >= len(s) {
		return
	}
	for i := index + 1; i <= len(s); i++ {
		curr := s[index:i]
		if isPalindrome(curr) {
			temp = append(temp, curr)
			helper(s, i, temp, used+i-index, res)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func copyArray(arr []string)[]string{
    res := make([]string, len(arr))
    for i := range len(arr){
        res[i] = arr[i]
    }
    return res
}