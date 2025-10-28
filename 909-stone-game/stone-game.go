func stoneGame(piles []int) bool {
    dp := make([][]int, len(piles))
    for i := range len(dp){
        dp[i] = make([]int, len(piles))
    }
    for i := range len(dp){
        for j := range len(dp[0]){
            dp[i][j] = -1
        }
    }
	xx := helper(piles, 0, len(piles)-1, dp)
	return xx > 0
}

func helper(arr []int, left int, right int, dp [][]int) int {
	if left > right {
		return 0
	}
    if dp[left][right] != -1{
        return dp[left][right]
    }
	x := arr[left] - helper(arr, left+1, right, dp)
	y := arr[right] - helper(arr, left, right-1, dp)
    dp[left][right] = max(x, y)
 	return max(x, y)
}