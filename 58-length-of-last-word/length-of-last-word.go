func lengthOfLastWord(s string) int {
	n := len(s) - 1
	for n >= 0 {
		if s[n] == 32 {
			n -= 1
		} else {
			break
		}
	}
	m := n
	for m >= 0 {
		if s[m] == 32 {
			break
		} else {
			m -= 1
		}
	}
	return n - m
}

