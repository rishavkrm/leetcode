func gameOfLife(board [][]int) {
	for i := range len(board) {
		for j := range len(board[0]) {
			nextState(&board, i, j)
		}
	}
	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == 5 {
				board[i][j] = 1
			}
			if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

func nextState(board *[][]int, i int, j int) {
	n := [][]int{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		{i, j - 1}, {i, j + 1},
		{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}}
	count := 0
	for _, x := range n {
		p, q := x[0], x[1]
		if p < 0 || p >= len(*board) || q < 0 || q >= len((*board)[0]) {
			continue
		}
		if (*board)[p][q] == 1 || (*board)[p][q] == 3 {
			count += 1
		}
	}
	// now alive next dead = 3
	if (*board)[i][j] == 1 {
		if count != 2 && count != 3 {
			(*board)[i][j] = 3
		}
	} else {
		if count == 3 {
			(*board)[i][j] = 5
		}
	}

}