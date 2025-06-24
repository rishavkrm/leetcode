func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	degrees := make(map[int]int)
	ts := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
		degrees[i] = 0
	}
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
		degrees[prerequisites[i][1]] += 1
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if degrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		arr := graph[queue[0]]
		for j := 0; j < len(arr); j++ {
			degrees[arr[j]] -= 1
			if degrees[arr[j]] == 0 {
				queue = append(queue, arr[j])
			}
		}
		ts = append(ts, queue[0])
		queue = queue[1:]

	}
	fmt.Println(ts)
	if len(ts) < numCourses {
		return false
	}
	return true

}