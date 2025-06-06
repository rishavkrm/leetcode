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