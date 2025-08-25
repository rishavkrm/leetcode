func generateParenthesis(n int) []string {
	result := []string{}
	helper(n, "", &result)
	return result
}

func helper(n int, curr string, result *[]string) {
	if len(curr) == 2*n {
		if validParantheses(curr) {
			*result = append(*result, curr)
		}
		return
	}
	options := []string{"(", ")"}
	for _, option := range options {
		curr += option
		helper(n, curr, result)
		curr = curr[:len(curr)-1]
	}
}

func validParantheses(s string) bool {
	d := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []byte{}
	for i, _ := range s {
		c := s[i]
		_, exists := d[c]
		if exists {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			popped := stack[len(stack)-1]
			if d[popped] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}