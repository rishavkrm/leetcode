func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := m - 1
	for i <= j {
		mid := (i + j) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
    if (i !=0 && matrix[i-1][n-1] >= target) || (i >= m){
        i = i-1
    }
    fmt.Println(i)
	arr := matrix[i]
	i = 0
	j = n-1
	for i <= j {
		mid := (i + j) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
    return false

}