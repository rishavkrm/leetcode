func isValid(s string) bool {
	d := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte
	for i := range len(s) {
		_, exists := d[s[i]]
		if exists {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			popped := stack[len(stack)-1]
			if d[popped] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}
	if len(stack) == 0 {
		return true
	}
	return false

}