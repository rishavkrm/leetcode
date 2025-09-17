func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	hash := map[string]bool{}
	helper(nums, 0, -200, []int{}, &result, hash)
	return result
}

func helper(nums []int, i int, currMax int, temp []int, result *[][]int, hash map[string]bool) {
	if i >= len(nums) {
		if len(temp) >= 2 {
            curr := createString(temp)
			_, ok := hash[curr]
			if !ok {
				hash[curr] = true
				(*result) = append(*result, copyArray(temp))
			}
		}
		return
	}
	if nums[i] >= currMax {
		temp = append(temp, nums[i])
		helper(nums, i+1, nums[i], temp, result, hash)
		temp = temp[:len(temp)-1]
		helper(nums, i+1, currMax, temp, result, hash)
	} else {
		helper(nums, i+1, currMax, temp, result, hash)
		temp = append([]int{}, nums[i])
		helper(nums, i+1, nums[i], temp, result, hash)
	}
}

func createString(nums []int) string {
	res := ""
	for i := range len(nums) {
		res += fmt.Sprintf("%d,", nums[i])
	}
	return res
}

func copyArray(nums []int) []int {
	res := make([]int, len(nums))
	for i := range len(nums) {
		res[i] = nums[i]
	}
	return res
}