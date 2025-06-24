func longestIncreasingPath(matrix [][]int) int {
	maxLength := 0
	cache := make([][]int, len(matrix))
	for i := range len(cache) {
		cache[i] = make([]int, len(matrix[0]))
	}

	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			maxLength = max(maxLength, helper(matrix, cache, i, j, -1))
		}
	}
	return maxLength
}
func helper(matrix, cache [][]int, row, col int, currMax int) int {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
		return 0
	}
	if matrix[row][col] <= currMax {
		return 0
	}
    if cache[row][col] != 0{
        return cache[row][col]
    }
	currMax = matrix[row][col]
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	maxLength := 0
	for i := range len(directions) {
		d := directions[i]
		x := helper(matrix, cache, row+d[0], col+d[1], currMax)
		maxLength = max(maxLength, x)
	}
	cache[row][col] = maxLength + 1
	return maxLength + 1
}
