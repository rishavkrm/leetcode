func exist(board [][]byte, word string) bool {
	heads := make([][]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				heads = append(heads, []int{i, j})
			}
		}
	}
	visited := make([][]bool, len(board))

	for i := range heads {
		for i := 0; i < len(board); i++ {
			visited[i] = make([]bool, len(board[0]))
		}
		res := dfs(board, word, visited, heads[i][0], heads[i][1], "")
		if res == true {
			return res
		}
	}

	return false

}
func dfs(board [][]byte, word string, visited [][]bool, row, col int, curr string) bool {
	if row >= len(board) || row < 0 || col >= len(board[0]) || col < 0 {
		return false
	}
	if !visited[row][col] {
		visited[row][col] = true
		curr += string(board[row][col])

		if curr == word {
			return true
		}
		if len(curr) >= len(word) {
			visited[row][col] = false
			return false
		}
		if word[:len(curr)] != curr {
			visited[row][col] = false
			return false
		}
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for i := range len(directions) {
			d := directions[i]
			res := dfs(board, word, visited, row+d[0], col+d[1], curr)
			if res == true {
				return true
			}
		}
		visited[row][col] = false
		return false
	}
	return false

}

// a, b
// c, d