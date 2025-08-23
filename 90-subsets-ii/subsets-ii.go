func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	result := make([][]int, 0)
	curr := make([]int, 0)
	helper(nums, 0, curr, &result)
	return result
}

func helper(nums []int, index int, curr []int, result *[][]int) {
	fmt.Println(curr)
	*result = append(*result, copyArray(curr))
	for i := index; i < len(nums); i++ {
		if i > index && nums[i-1] == nums[i] {
			continue
		}
		curr = append(curr, nums[i])
		helper(nums, i+1, curr, result)
		curr = curr[:len(curr)-1]
	}
}

func copyArray(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i])
	}
	return res
}