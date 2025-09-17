func isAdditiveNumber(num string) bool {
	return helper(num, 0, "", []string{})
}

func helper(num string, index int, curr string, temp []string) bool {
	if index >= len(num) {
		if len(temp) > 2 {
			fmt.Println(temp)
			return true
		}
		return false
	}
	res := false
	for i := index; i < len(num); i++ {
		str := num[index : i+1]
		if len(str) > 1 {
			if str[0] == '0' {
				continue
			}
		}
		if len(temp) < 2 {
			temp = append(temp, str)
			res = res || helper(num, i+1, curr+str, temp)
			temp = temp[:len(temp)-1]
		} else {
			n, _ := strconv.Atoi(str)
			n1, _ := strconv.Atoi(temp[len(temp)-1])
			n2, _ := strconv.Atoi(temp[len(temp)-2])
			if n == n1+n2 {
				temp = append(temp, str)
				res = res || helper(num, i+1, curr+str, temp)
				temp = temp[:len(temp)-1]
			}
		}
        if res{
            return res
        }

	}
    return res
}