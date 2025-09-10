func isBipartite(graph [][]int) bool {
	if len(graph) == 1 {
		return true
	}
	parent := make([]int, len(graph))
	rank := make([]int, len(graph))
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	for node, edges := range graph {
		for _, edge := range edges {
			nbrEdges := graph[edge]
			for _, x := range nbrEdges {
				union(node, x, &parent, &rank)
			}
		}
	}
	for node, edges := range graph {
		for _, edge := range edges {
			if find(node, &parent) == find(edge, &parent){
                return false
            }
		}
	}
	return true
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