func reverseWords(s string) string {
	res := ""
	n := len(s) - 1
	i := n
	for i > -1 {
		for i > -1 && s[i] == 32 {
			i -= 1
		}
		m := i
		for i > -1 && s[i] != 32 {
			i -= 1
		}
		if m > i {
			res += s[i+1 : m+1]
		}
		res += " "
	}
	k := len(res) - 1
	for k >= 0 {
		if res[k] == 32 {
			k -= 1
		} else {
			break
		}
	}
	return res[0 : k+1]
}