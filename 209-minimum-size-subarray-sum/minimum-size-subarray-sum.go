// Start with a minLength of len(array) + 1, and two pointers start and end
// maintain a currSum
// if currSum >= target, then update the minLength and decrease the start
// otherwise (currSum < target), then increase end and continue the process

func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	start := 0
	end := 0
	currSum := nums[0]
	for end < len(nums) {
		if currSum >= target {
			minLength = min(minLength, end-start+1)
			currSum -= nums[start]
			start += 1
		} else {
			if end+1 < len(nums) {
				currSum += nums[end+1]
			}
			end += 1
		}

	}

	if minLength > len(nums) {
		return 0
	}
	return minLength
}