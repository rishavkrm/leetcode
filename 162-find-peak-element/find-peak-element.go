func findPeakElement(nums []int) int {
	i := 0
	j := len(nums) - 1
	if len(nums) == 1 {
		return 0
	}
	for i <= j {
		mid := (i + j) / 2
		if isPeak(nums, mid) {
			return mid
		} else if nums[mid+1] > nums[mid] {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

func isPeak(nums []int, index int) bool {
	if index == 0 {
		if nums[index+1] < nums[index] {
			return true
		} else {
			return false
		}
	}
	if index == len(nums)-1 {
		if nums[index-1] < nums[index] {
			return true
		} else {
			return false
		}
	}
	if nums[index] > nums[index+1] && nums[index] > nums[index-1] {
		return true
	}
	return false
}