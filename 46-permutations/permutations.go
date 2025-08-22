func permute(nums []int) [][]int {
	result := make([][]int, 0)
	isUsed := make([]bool, len(nums))
	helper(nums, isUsed, []int{}, &result)
	return result
}

func helper(nums []int, isUsed []bool, curr []int, result *[][]int) {
	if len(curr) == len(nums) {
		*result = append(*result, copyArray(curr))
		return
	}
	for i := 0; i < len(nums); i++ {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		curr = append(curr, nums[i])
		helper(nums, isUsed, curr, result)
		isUsed[i] = false
		curr = curr[:len(curr)-1]
	}
}

func copyArray(arr []int) []int {
	temp := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i])
	}
	return temp
}
