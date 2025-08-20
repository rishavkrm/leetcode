func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	i := 0
	j := len(nums) - 1
	if nums[i] < nums[j] {
		return nums[0]
	}
	for i <= j {
		mid := (i + j) / 2
		if isMinIndex(nums, mid) {
			return nums[mid]
		} else if nums[mid] >= nums[0] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
    return nums[0]
}

func isMinIndex(nums []int, index int) bool {
	if index == 0 {
		return false
	}
	if nums[index-1] > nums[index] {
		return true
	}
	return false
}