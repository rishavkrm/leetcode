func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	for i, _ := range parent {
		parent[i] = i
	}
	options := [][]int{{1, 0}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for _, option := range options {
					x := i + option[0]
					y := j + option[1]
					if x < m && y < n && grid[x][y] == 1 {
                        // fmt.Println(x, y, i, j)
						union(findIndex(x, y, m, n), findIndex(i, j, m, n), &parent, &rank)
					}
				}
			}
		}
	}
	parents := map[int]int{}
	maxi := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				parents[find(findIndex(i, j, m, n), &parent)] += 1
				maxi = max(maxi, parents[find(findIndex(i, j, m, n), &parent)])
			}
		}
	}
    fmt.Println(parents)
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

func findIndex(i int, j int, m int, n int) int {
	return j + (i)*n
}
