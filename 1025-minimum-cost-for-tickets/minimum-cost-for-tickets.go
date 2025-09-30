func mincostTickets(days []int, costs []int) int {
    dp := make([]int, len(days))
    for i := range len(dp){
        dp[i] = -1
    }
    return helper(days, costs, 0, dp)
}

func helper(days []int, costs []int, index int, dp []int) int {
	if index >= len(days) {
		return 0
	}
	x1, x7, x30 := costs[0], costs[1], costs[2]
    z1,z7,z30 := index, index, index
    if dp[index] != -1{
        return dp[index]
    }
    for i := index+1; i < len(days); i++{
        if days[i] < days[index] + 7{
            z7 += 1
        }
        if days[i] < days[index] + 30{
            z30 += 1
        }
    }    
    y1  := x1  + helper(days, costs, z1+1, dp)
    y7  := x7  + helper(days, costs, z7+1, dp)
    y30 := x30 + helper(days, costs, z30+1, dp)
    dp[index] = min(min(y1,y7),y30)
    return min(min(y1,y7),y30)
}