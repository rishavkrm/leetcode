// [["A","B","C","E"],
// ["S","F","E","S"],
// ["A","D","E","E"]]

func exist(board [][]byte, word string) bool {
	for i := range len(board) {
		for j := range len(board[0]) {
			if helper(board, word, i, j, 0,getVisited(board)) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, word string, i int, j int, k int, visited *[][]bool) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	if (*visited)[i][j] {
		return false
	}
	(*visited)[i][j] = true
	options := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, option := range options {
		if helper(board, word, i+option[0], j+option[1], k+1, visited) {
			return true
		}

	}
	(*visited)[i][j] = false
	return false
}

func getVisited(board [][]byte) *[][]bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	return &visited
}