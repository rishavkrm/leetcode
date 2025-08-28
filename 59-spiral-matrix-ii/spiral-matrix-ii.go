func generateMatrix(n int) [][]int {
    res := make([][]int, n)
    for i := range(len(res)){
        res[i] = make([]int, n)
    }
	cRows, cCols := map[int]bool{}, map[int]bool{}
	helper(n, 0, 0, cRows, cCols, "right", 1, &res)
	return res

}

func helper(n int, i int, j int, cRows map[int]bool, cCols map[int]bool, dir string, count int, res *[][]int) {
	if count > n*n {
		return
	}
	(*res)[i][j] = count
	next := getNext([]int{i, j}, dir)
	if dir == "right" && (next[1] >= n || cCols[next[1]]) {
		cRows[i] = true
		next = []int{i + 1, j}
		dir = "bottom"
	}
	if dir == "left" && (next[1] < 0 || cCols[next[1]]) {
		cRows[i] = true
		next = []int{i - 1, j}
		dir = "top"
	}
	if dir == "top" && (next[0] < 0 || cRows[next[0]]) {
		cCols[j] = true
		next = []int{i, j + 1}
		dir = "right"
	}
	if dir == "bottom" && (next[0] >= n || cRows[next[0]]) {
		cCols[j] = true
		next = []int{i, j - 1}
		dir = "left"
	}
	helper(n, next[0], next[1], cRows, cCols, dir, count+1, res)

}

func getNext(next []int, dir string) []int {
	if dir == "right" {
		next[1] += 1
	}
	if dir == "bottom" {
		next[0] += 1
	}
	if dir == "left" {
		next[1] -= 1
	}
	if dir == "top" {
		next[0] -= 1
	}
	return next
}