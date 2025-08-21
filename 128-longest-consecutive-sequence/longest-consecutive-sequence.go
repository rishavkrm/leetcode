func longestConsecutive(nums []int) int {
	hMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hMap[nums[i]] = 1
	}
	maxi := 0
	i := 0
for curr, _ := range hMap {
		if hMap[curr-1] == 0 {
			count := 1
			for hMap[curr+1] == 1 {
				count += 1
				curr += 1
			}
			maxi = max(maxi, count)
		}
		i += 1

	}
	return maxi

}