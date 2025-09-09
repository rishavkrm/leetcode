func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
		parent[i] = i
	}
	for _, edge := range edges {
        union(edge[0], edge[1], &parent, &rank)
	}
    return find(source, &parent) == find(destination, &parent)
	
}

func find(k int, parent *[]int) int {
	if !((*parent)[k] == k) {
		(*parent)[k] = find((*parent)[k], parent)
	}
	return (*parent)[k]
}
func union(x1 int, x2 int, p *[]int, r *[]int) {
	x1Parent := find(x1, p)
	x2Parent := find(x2, p)
	if x1Parent == x2Parent {
		return
	}
	if (*r)[x1Parent] > (*r)[x2Parent] {
		(*p)[x2Parent] = x1Parent
	} else if (*r)[x1Parent] < (*r)[x2Parent] {
		(*p)[x1Parent] = x2Parent
	} else {
		(*p)[x1Parent] = x2Parent
		(*r)[x2Parent] += 1
	}
}