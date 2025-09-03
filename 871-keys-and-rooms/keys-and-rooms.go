func canVisitAllRooms(rooms [][]int) bool {
	return dfs(rooms)

}

func dfs(graph [][]int) bool {
	visited := map[int]bool{}
	stack := []int{0}
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[popped] = true
		for _, nbr := range graph[popped] {
			if !visited[nbr] {
				stack = append(stack, nbr)
			}
		}
	}
	if len(visited) == len(graph) {
		return true
	}
	return false
}