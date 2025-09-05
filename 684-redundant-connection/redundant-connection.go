func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
    var result []int
	for _, edge := range edges {
        if union(edge[0]-1, edge[1]-1, &parent, &rank){
            result = edge
        }
	}
    return result
}

func find(k int, parent *[]int) int {
	if !((*parent)[k] == k) {
		(*parent)[k] = find((*parent)[k], parent)
	}
	return (*parent)[k]
}

func union(x int, y int, parent *[]int, rank *[]int) bool {
	pX := find(x, parent)
	pY := find(y, parent)
	if pX == pY {
		return true
	}
	if (*rank)[pX] > (*rank)[pY] {
		(*parent)[pY] = pX
	} else if (*rank)[pY] > (*rank)[pX] {
		(*parent)[pX] = pY
	} else {
		(*parent)[pX] = pY
		(*rank)[pY] += 1
	}
	return false
}