func minimumTotal(triangle [][]int) int {
    dp := make([][]int, len(triangle))
    for i := range len(dp){
        dp[i] = make([]int, len(dp))
    }
    for i := range len(dp){
        for j := range len(dp[0]){
            dp[i][j] = -1001
        }
    }
    return helper(triangle, 0, 0, &dp)    
}

func helper(nums [][]int, i int, j int, dp *[][]int)int{
    if i >= len(nums){
        return 0
    }
    if j >= len(nums[i]){
        return 0
    }
    if (*dp)[i][j] != -1001{
        return (*dp)[i][j]
    }
    // case 1 
    x := nums[i][j] + helper(nums, i+1, j+1, dp)
    // case 2
    y := nums[i][j] + helper(nums, i+1, j, dp)
    (*dp)[i][j] = min(x, y)
    return min(x, y)
}
