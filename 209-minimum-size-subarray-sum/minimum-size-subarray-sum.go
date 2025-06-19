func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	start := 0
	end := 0
	currSum := nums[0]
	for end < len(nums) {
		if currSum == target {
			minLength = min(minLength, end-start+1)
			currSum -= nums[start]
			start += 1
		} else if currSum < target {
			if end+1 < len(nums) {
				currSum += nums[end+1]
			}
			end += 1
		} else {
			minLength = min(minLength, end-start+1)
			currSum -= nums[start]
			start += 1

		}

	}

	if minLength > len(nums) {
		return 0
	}
	return minLength
}

// 2, 5, 6, 8, 12, 15
// 