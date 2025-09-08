func removeStones(stones [][]int) int {
	n := len(stones)
    if n == 1{
        return 0
    }
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			row1 := stones[i][0]
			row2 := stones[j][0]
			col1 := stones[i][1]
			col2 := stones[j][1]
			if (row1 == row2) || (col1 == col2) {
				union(i, j, &parent, &rank)
			}
		}
	}
    maxi := 0
    for i := 0; i < len(parent); i++{
        if find(i, &parent) != i{
            maxi += 1
        }
    }
    return maxi

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
