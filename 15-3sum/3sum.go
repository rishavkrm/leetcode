func threeSum(nums []int) [][]int {
	target := 0
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	res := make([][]int, 0)
	i := 0
	for i < len(nums) {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right -= 1
			} else if sum < target {
				left += 1
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// right -= 1
				y := nums[left]
				for left < right && nums[left] == y {
					left += 1
				}
			}
		}
        y := nums[i]
        for i < len(nums) && nums[i] == y {
			i += 1
        } 

	}
	return res
}