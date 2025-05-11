func coinChange(coins []int, amount int) int {
    dp := make([][]int, amount+1)
    for i := range(len(dp)){
        dp[i] = make([]int, len(coins))
    }
    for i := range(len(dp)){
        for j := range(len(dp[0])){
            dp[i][j] = -2
        }
    }
    res := coinChangeHelper(coins, amount, 0, dp)
    if res == 10001{
        return -1
    }
    return res
}

func coinChangeHelper(coins []int, amount int, index int, dp [][]int)int{
    if amount == 0{
        return 0
    }
    if amount < 0 || index >= len(coins){
        return 10001
    }
    if dp[amount][index] != -2{
        return dp[amount][index]
    }
    // Unbounded knapsack
    // take and potentially take it again, that is why not increasing index
    take := 1 + coinChangeHelper(coins, amount - coins[index], index, dp)
    not_take := coinChangeHelper(coins, amount, index+1, dp)
    dp[amount][index] =  Min(take, not_take)
    return dp[amount][index]
}

func Min(a int, b int) int{
    if a > b{
        return b
    }
    return a
}