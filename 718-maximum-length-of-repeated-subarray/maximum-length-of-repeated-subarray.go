func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range len(dp) {
		dp[i] = make([]int, len(nums2))
	}
	for i := range len(dp) {
		for j := range len(dp[0]) {
			dp[i][j] = -1
		}
	}
	currMax := 0
	x := helper(nums1, nums2, 0, 0, &currMax, dp)
    // fmt.Println(dp)
    return x
}

func helper(nums1 []int, nums2 []int, i1 int, i2 int, currMax *int, dp [][]int) int {
	if i1 >= len(nums1) || i2 >= len(nums2) {
		return *currMax
	}
	if dp[i1][i2] != -1 {
		return dp[i1][i2]
	}
	x, y, z := 0, 0, 0
	if nums1[i1] == nums2[i2] {
		count := 1
		for i1+count < len(nums1) && i2+count < len(nums2) && nums1[i1+count] == nums2[i2+count] {
			count += 1
		}
        *currMax = max(*currMax, count)
		x = helper(nums1, nums2, i1+count, i2+count, currMax, dp)
	}
	y = helper(nums1, nums2, i1+1, i2, currMax, dp)
	z = helper(nums1, nums2, i1, i2+1, currMax, dp)
	dp[i1][i2] = max(z, max(x, y))
	return max(z, max(x, y))
}