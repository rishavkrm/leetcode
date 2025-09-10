func possibleBipartition(n int, dislikes [][]int) bool {
	graph := map[int][]int{}
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		graph[i] = []int{}
		parent[i] = i
	}
	for _, xx := range dislikes {
		graph[xx[0]-1] = append(graph[xx[0]-1], xx[1]-1)
		graph[xx[1]-1] = append(graph[xx[1]-1], xx[0]-1)
	}

	for i := range len(graph) {
		for _, nbr := range graph[i] {
			for _, nbr2 := range graph[nbr] {
				union(i, nbr2, &parent, &rank)
			}
		}
	}
	for i := range len(graph) {
		for _, nbr := range graph[i] {
			if find(i, &parent) == find(nbr, &parent) {
				return false
			}
		}
	}
    fmt.Println(graph)
	fmt.Println(parent)

	return true
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