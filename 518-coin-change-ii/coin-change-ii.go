func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i := range len(dp) {
		dp[i] = make([]int, amount+1)
		for j := range len(dp[i]) {
			dp[i][j] = -1
		}
	}

	result := helper(coins, amount, 0, &dp)
	return result
}

func helper(coins []int, amount int, i int, dp *[][]int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 || i >= len(coins) {
		return 0
	}
	if (*dp)[i][amount] != -1 {
		return (*dp)[i][amount]
	}
	x := helper(coins, amount-coins[i], i, dp)
	y := helper(coins, amount, i+1, dp)
	z := x + y
	(*dp)[i][amount] = z
	return z
}