func search(nums []int, target int) bool {
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	i := 0
	j := len(nums) - 1
	// 1f 1st element is lesser than last element that means no rotation
	if nums[i] < nums[j] {
		return binarySearch(nums, target, i, j)
	}
	// if we draw a graph it will be like a increasing line then dip at pivvot index
	// and then inceasing again but it wont even cross the 0th element again
	// let's find pivvot index now
	pivvotIndex := -1
	for i <= j {
		mid := (i + j) / 2
		if isPivvotIndex(nums, mid) {
			pivvotIndex = mid
			break
		}
		if nums[mid] >= nums[0] {
			// left half
			i = mid + 1
		} else if nums[mid] < nums[0] {
			// right half
			j = mid - 1
		}
	}
	if pivvotIndex == -1 {
		pivvotIndex = findMinIndex(nums)
	}

	i = 0
	j = len(nums) - 1
	if pivvotIndex == 0 {
		return binarySearch(nums, target, i, j)
	}
	if target >= nums[0] {
		j = pivvotIndex - 1
	} else {
		i = pivvotIndex
	}
	return binarySearch(nums, target, i, j)

}

func isPivvotIndex(nums []int, index int) bool {
	if index == 0 {
		return false
	}
	if nums[index-1] > nums[index] {
		return true
	}
	return false
}

func binarySearch(nums []int, target int, i int, j int) bool {
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

func findMinIndex(nums []int) int {
	min := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			min = i
		}
	}
	return min
}