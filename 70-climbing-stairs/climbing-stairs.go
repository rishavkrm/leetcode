func climbStairs(n int) int {
    dp := make([]int, n+1)
    return Temp(n, n, dp)
}

func Temp(curr int, n int, dp []int) int{
    if curr <= 0 {
        return 0
    }
    if curr == 1{
        return 1
    }
    if curr == 2{
        return 2
    }
    if dp[curr] != 0 {
        return dp[curr]
    }
    dp[curr] =  Temp(curr-1,n, dp) + Temp(curr-2, n, dp)
    return dp[curr]
}