func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	helper(k, n, 0, []int{}, 0, &res)
	return res
}

func helper(k int, n int, index int, curr []int, sum int, res *[][]int) {
	if len(curr) == k {
		if sum == n {
            *res = append(*res, copyArray(curr))
		}
		return
	}
	if index >= 9 {
		return
	}
	for i := index; i < 9; i++ {
		curr = append(curr, i+1)
		helper(k, n, i+1, curr, sum+i+1, res)
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