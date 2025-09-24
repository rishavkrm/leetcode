func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range len(dp) {
		dp[i] = make([]int, amount+1)
	}

	result := helper(coins, amount, 0, &dp)
	if result == 99999 {
		return -1
	}
	return result
}

func helper(coins []int, amount int, index int, dp *[][]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 || index >= len(coins) {
		return 99999
	}
	if (*dp)[index][amount] != 0 {
		return (*dp)[index][amount]
	}
	mini := 99999
	for i := index; i < len(coins); i++ {
		x := 1 + helper(coins, amount-coins[i], i,dp)
		mini = min(mini, x)
	}
	(*dp)[index][amount] = mini
	return mini
}