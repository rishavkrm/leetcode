func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	rank := make([]int, m*n)
	parent := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	options := [][]int{{1, 0}, {0, 1}}
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == '1' {
				for _, o := range options {
					x := o[0] + i
					y := o[1] + j
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
						index1 := getIndex(x, y, m, n)
						index2 := getIndex(i, j, m, n)
						union(index1, index2, &parent, &rank)
					}
				}
			}
		}
	}
	mp := map[int]int{}
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == '1' {
				index := getIndex(i, j, m, n)
				p := find(parent[index], &parent)
				mp[p] += 1
			}
		}
	}
	return len(mp)
}

func find(x int, p *[]int) int {
	if !((*p)[x] == x) {
		(*p)[x] = find((*p)[x], p)
	}
	return (*p)[x]
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

func getIndex(i int, j int, m int, n int) int {
	return j + (i)*n
}

// [["1","0","1","1","1"],
//  ["1","0","1","0","1"],
//  ["1","1","1","0","1"]]

// 5 1 2 2 2
// 5 6 2 8 2
// 5 5 2 13 2

