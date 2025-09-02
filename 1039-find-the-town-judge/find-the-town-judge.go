func findJudge(n int, trust [][]int) int {
	graph1 := map[int][]int{}
	graph2 := map[int][]int{}
	for i := range n {
		graph1[i+1] = []int{}
		graph2[i+1] = []int{}
	}
	for _, t := range trust {
		graph1[t[0]] = append(graph1[t[0]], t[1])
		graph2[t[1]] = append(graph2[t[1]], t[0])
	}
	fmt.Println(graph1)
	fmt.Println(graph2)
	for i := range n + 1 {
		if i == 0 {
			continue
		}
		if len(graph2[i]) == n-1 && len(graph1[i]) == 0 {
			return i
		}
	}
	return -1

}