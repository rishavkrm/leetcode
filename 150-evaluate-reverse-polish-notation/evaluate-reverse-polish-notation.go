func evalRPN(tokens []string) int {
	stack := []string{}
	operands := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	for i := range len(tokens) {
		c := tokens[i]
		if operands[c] {
			stack = handleCalculation(stack, c)
		} else {
			stack = append(stack, c)
		}
	}
    res,_ := strconv.Atoi(stack[0])
	return res
}

func handleCalculation(stack []string, operand string) ([]string) {
	c1, _ := strconv.Atoi(stack[len(stack)-2])
	c2, _ := strconv.Atoi(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	var result int
	if operand == "*" {
		result = c1 * c2
	} else if operand == "/" {
		result = c1 / c2
	} else if operand == "-" {
		result = c1 - c2
	} else {
		result = c1 + c2
	}
	stack[len(stack)-1] = fmt.Sprintf("%d", result)
	return stack
}