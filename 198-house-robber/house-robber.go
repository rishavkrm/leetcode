func rob(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++{
        dp[i] = -1
    }
    if len(nums) == 1{
        return nums[0]
    }
    xx := nums[0] + Temp(nums, 0, dp)
    yy := nums[1] + Temp(nums, 1, dp)
    return Max(xx, yy)
}

func Temp(nums []int, index int, dp []int) int{
    if index >= len(nums){
        return 0
    }
    if dp[index] != -1{
        return dp[index]
    }
    yy := 0
    xx := 0
    if index + 3 < len(nums){
        yy = nums[index+3] + Temp(nums, index+3, dp)
    }
    if index + 2 < len(nums){
        xx = nums[index+2] + Temp(nums, index+2, dp)
    }
    dp[index] =  Max(xx, yy)
    return dp[index]
}

func Max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}