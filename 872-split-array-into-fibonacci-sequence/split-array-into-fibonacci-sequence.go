func splitIntoFibonacci(num string) []int {
	return helper(num, 0, []int{}, 0)
}

func helper(num string, index int, temp []int, used int) []int {
	if len(temp)>2 && used == len(num) {
		return temp
	}
	if index > len(num) {
		return []int{}
	}
	for i := index + 1; i <= len(num); i++ {
		curr := num[index:i]
        if len(curr) > 1 && curr[0] == '0'{
            continue
        }
        currInt,_ := strconv.Atoi(curr)
        if currInt > 2147483648{
            continue
        }
		temp = append(temp, currInt)
        if passes(temp){
            res := helper(num, i, temp, used+len(curr))
            if len(res) > 0 {
                return res
            }
        }
        temp = temp[:len(temp)-1]
	}
	return []int{}
}

func passes(arr []int) bool {
	if len(arr) < 3 {
		return true
	}
	i := len(arr)-1
	x := arr[i-2]
	y := arr[i-1]
	z := arr[i]
	if x+y != z {
		return false
	}
	return true
}