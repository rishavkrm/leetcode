func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
	helper(candidates, target, []int{}, 0, &result)
	return result
}
func helper(nums []int, target int, curr []int, index int, result *[][]int) {
	if target == 0 {
		*result = append(*result, copyArray(curr))
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(nums); i++ {
		target -= nums[i]
		curr = append(curr, nums[i])
		helper(nums, target, curr, i, result)
		target += nums[i]
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
