func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

// 2x+1/2 = x + 0.5