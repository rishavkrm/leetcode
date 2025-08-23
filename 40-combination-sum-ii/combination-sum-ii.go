func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	result := make([][]int, 0)
	m := map[string]int{}
	helper(candidates, target, []int{}, 0, &result, m)
	return result
}
func helper(nums []int, target int, curr []int, index int, result *[][]int, m map[string]int) {
	if target == 0 {
		*result = append(*result, copyArray(curr))
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(nums); i++ {
        if i > index && nums[i-1] == nums[i]{
            continue
        }
		target -= nums[i]
		curr = append(curr, nums[i])
		helper(nums, target, curr, i+1, result, m)
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

// func convertToString(arr []int) string {
// 	s := ""
// 	for i := 0; i < len(arr); i++ {
// 		s += fmt.Sprintf("%d", arr[i])
// 		if i != len(arr)-1 {
// 			s += ","
// 		}
// 	}
// 	return s
// }