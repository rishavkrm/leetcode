func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	res := []int{-1, -1}
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != nums[mid]) {
			res[0] = mid
            break
		} else if nums[mid] >= target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	i = 0
	j = len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] > nums[mid]) {
			res[1] = mid
            break
		} else if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return res
}