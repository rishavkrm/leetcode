func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int, 0)
	degrees := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		degrees[i] = 0
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
		degrees[prerequisites[i][1]] += 1
	}
	queue := make([]int, 0)
	for i := 0; i < len(degrees); i++ {
		if degrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	topSort := make([]int, numCourses)
	index := numCourses - 1
	for len(queue) > 0 {
		topSort[index] = queue[0]
		neighbours := graph[queue[0]]
		for i := range len(neighbours) {
			degrees[neighbours[i]] -= 1
			if degrees[neighbours[i]] == 0 {
				queue = append(queue, neighbours[i])
			}
		}
		index -= 1
		queue = queue[1:]
	}
	if index != -1 {
		return []int{}
	}
	return topSort

}   