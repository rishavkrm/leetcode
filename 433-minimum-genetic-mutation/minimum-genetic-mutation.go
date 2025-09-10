func minMutation(startGene string, endGene string, bank []string) int {
	if !slices.Contains(bank, startGene) {
		bank = append(bank, startGene)
	}
	if !slices.Contains(bank, endGene) {
		return -1
	}
	hashMap := map[string]int{}
	for i, gene := range bank {
		hashMap[gene] = i
	}
	graph := map[int][]int{}
	for i := range len(bank) {
		graph[i] = []int{}
	}

	for ii, gene := range bank {
		options := []byte{'A', 'C', 'G', 'T'}
		for i, x := range gene {
			for _, option := range options {
				if byte(x) == option {
					continue
				}
				temp := []byte(gene)
				temp[i] = option
				_, exists := hashMap[string(temp)]
				if exists {
					graph[ii] = append(graph[ii], hashMap[string(temp)])
				}
			}
		}
	}
	fmt.Println(graph)
	return bfs(graph, hashMap[startGene], hashMap[endGene])
}

func bfs(graph map[int][]int, s int, end int) int {
	queue := []int{s}
	dist := map[int]int{}
	visited := map[int]bool{}
	dist[s] = 0
	visited[s] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, nbr := range graph[curr] {
			if !visited[nbr] {
				visited[nbr] = true
				dist[nbr] = dist[curr] + 1
				queue = append(queue, nbr)
			}
		}
	}
	val, exists := dist[end]
	if exists {
		return val
	}
	return -1
}
