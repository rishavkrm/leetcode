func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	blackListed := map[int]bool{}
	for i := 0; i < m*n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	options := [][]int{{0, 1}, {1, 0}}
	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				for _, option := range options {
					x := i + option[0]
					y := j + option[1]
					if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
						union(getIndex(i, j, m, n), getIndex(x, y, m, n), &parent, &rank)
					}
				}
			}
		}
	}
	for i := range m {
		for j := range n {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				blackListed[find(getIndex(i, j, m, n), &parent)] = true
			}
		}
	}
	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				if !blackListed[find(getIndex(i, j, m, n), &parent)] {
					board[i][j] = 'X'
				}
			}
		}
	}

}

// ["X","O","X"],
// ["O","X","O"],
// ["X","O","X"]

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

func getIndex(i int, j int, m int, n int) int {
	return j + (i)*n
}
