func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			s = s[i+1:]
		} else {
			break
		}
	}
	if len(s) == 0 {
		return 0
	}
    	j := len(s) - 1

	for j > 0 {
		if s[j] == ' ' {
			s = s[:j]
            j -= 1
		} else {
			break
		}
	}
	if len(s) == 0 {
		return 0
	}

	count := 1
	i = 0
	for i < len(s) {
		if s[i] == ' ' {
			count += 1
			for i < len(s) && s[i] == ' ' {
				i += 1
			}
		} else {
			i += 1
		}
	}
	return count
}