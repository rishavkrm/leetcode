// 1. sort the array
// 2. use two pointers for left and right
// 3. iterate over the array
// 4. if curr_num + left + right < target, then left += 1
// 5. if curr_num + left + right > target, then right -= 1
// 6. else curr_num + left + right == target, then add into the array
// 7. increase left until it is not equal to current left, otherwise it will be generate duplicate pairs  

func threeSum(array []int) [][]int {
	res := make([][]int, 0)
	// sort the array
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	target := 0
	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1
		if i > 0 && array[i] == array[i-1] {
			continue
		}
		for left < right {
			sum := array[i] + array[left] + array[right]
			if sum > target {
				right -= 1
			} else if sum < target {
				left += 1
			} else {
				res = append(res, []int{array[i], array[left], array[right]})
                curr := array[left]
				for left < right && array[left] == curr{
                    left += 1
                } 
				right -= 1
			}
		}
	}
	return res
}