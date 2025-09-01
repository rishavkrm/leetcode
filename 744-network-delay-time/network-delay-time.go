func networkDelayTime(times [][]int, n int, k int) int {
	visited := map[int]bool{}
	minDistance := map[int]int{}
	graph := map[int][][]int{}
	for i := range n {
		graph[i+1] = [][]int{}
		minDistance[i+1] = math.MaxInt32
	}
	minDistance[k] = 0

	for i := range len(times) {
		currNode := times[i][0]
		graph[currNode] = append(graph[currNode], []int{times[i][1], times[i][2]})
	}
	return Dijkstra(k, n, k, visited, minDistance, graph)

}

func Dijkstra(node int, n int, k int, v map[int]bool, minDistance map[int]int, graph map[int][][]int) int {
	for len(v) != n {
		v[node] = true
		neighbours := graph[node]
		nextNode := -1
		minDist := math.MaxInt32
		for i := range neighbours {
			nN := neighbours[i][0]
			dist := neighbours[i][1]
			if v[nN] {
				continue
			}
			minDistance[nN] = min(minDistance[node]+dist, minDistance[nN])
		}
		for i := range n {
			if i+1 == k || v[i+1]{
				continue
			}
			if minDist > minDistance[i+1] {
				nextNode = i + 1
				minDist = minDistance[i+1]
			}
		}
		if nextNode == -1 {
			break
		} else {
			node = nextNode
		}
	}
	res := -1
	for _, val := range minDistance {
		res = max(res, val)
	}
	if res == math.MaxInt32 {
		res = -1
	}
	return res
}