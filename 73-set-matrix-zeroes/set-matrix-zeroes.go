func setZeroes(matrix [][]int) {
	//
	maxi := matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxi = max(maxi, matrix[i][j])
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				makeZero(i, j, &matrix, maxi)
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == maxi+1 {
				(*(&matrix))[i][j] = 0
			}
		}
	}
}

func makeZero(row, col int, matrix *[][]int, maxi int) {
	for i := 0; i < len((*matrix)[0]); i++ {
		if (*matrix)[row][i] != 0 {
			(*matrix)[row][i] = maxi + 1
		}
	}
	for i := 0; i < len(*matrix); i++ {
		if (*matrix)[i][col] != 0 {
			(*matrix)[i][col] = maxi + 1
		}
	}
}