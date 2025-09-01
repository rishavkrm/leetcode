func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][][]int{}
	dp := make([][]int, n)
	for i := range len(dp) {
		dp[i] = make([]int, n)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}
	for i := range n {
		graph[i] = [][]int{}
	}
	for i := range flights {
		from, to, cost := flights[i][0], flights[i][1], flights[i][2]
		graph[from] = append(graph[from], []int{to, cost})
	}
	res := helper(graph, k, src, dst, 0, dp)
	if res == math.MaxInt32 {
		res = -1
	}
	return res
}

func helper(graph map[int][][]int, k int, src int, dst int, curr int, dp [][]int) int {
	if src == dst {
		return 0
	}
	if k < 0 {
		return math.MaxInt32
	}

	if dp[src][k] > -1 {
		return dp[src][k]
	}
	nbrs := graph[src]
	mini := math.MaxInt32
	for _, nbr := range nbrs {
		res := helper(graph, k-1, nbr[0], dst, curr+nbr[1], dp)
		if res < math.MaxInt32 {
			mini = min(mini, res+nbr[1])
		}

	}
	dp[src][k] = mini
	return mini
}