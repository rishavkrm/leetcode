func search(nums []int, target int) int {
	x := findPivvotIndex(nums)
    if target == nums[0]{
        return 0
    }
	if target > nums[len(nums)-1] {
		return binarySearch(nums, target, 0, x)
	}
	return binarySearch(nums, target, x, len(nums)-1)
}

func findPivvotIndex(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left

}

func binarySearch(nums []int, target int, i int, j int) int {
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}
