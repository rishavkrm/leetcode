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

func helper(coins []int, amount int, i int, dp *[][]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 || i >= len(coins) {
		return 99999
	}
	if (*dp)[i][amount] != 0 {
		return (*dp)[i][amount]
	}
	x := 1 + helper(coins, amount-coins[i], i, dp)
	y := helper(coins, amount, i+1, dp)
    z := min(x, y)
	(*dp)[i][amount] = z
	return z
}