func allPathsSourceTarget(graph [][]int) [][]int {
    result := [][]int{}
    dfs(graph, 0, []int{0}, map[int]bool{}, &result)
	return result
}

func dfs(graph [][]int, curr int, temp []int, visited map[int]bool, result *[][]int) {
	if curr == len(graph)-1 {
		*result = append(*result, copyArray(temp))
		return
	}
    if len(visited) == len(graph){
        return 
    }
    visited[curr] = true
    for _, nbr := range graph[curr]{
        if !visited[nbr] {
            temp = append(temp, nbr)
            dfs(graph, nbr, temp, visited, result)
            temp = temp[:len(temp)-1]
        }
    }
    visited[curr] = false
}

func copyArray(temp []int)[]int{
    res := make([]int, len(temp))
    for i, val := range temp{
        res[i] = val
    }
    return res
}