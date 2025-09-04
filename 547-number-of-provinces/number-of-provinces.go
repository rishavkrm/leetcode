func findCircleNum(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	rank := make([]int, len(isConnected))
	for i, _ := range parent {
		parent[i] = i
		rank[i] = 1
	}
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected); j++ {
			if i == j {
				continue
			}
			if isConnected[i][j] == 1 {
				union(i, j, &rank, &parent)
				fmt.Println(parent)
			}
		}
	}
	count := 0
	for i, p := range parent {
		if p == i {
			count += 1
		}
	}
	return count

}

func union(i int, j int, rank *[]int, parent *[]int) {
	parentI := find(i, parent)
	parentJ := find(j, parent)
	if parentI == parentJ {
		return
	}
	if (*rank)[parentI] > (*rank)[parentJ] {
		(*parent)[parentJ] = parentI
	} else if (*rank)[parentI] < (*rank)[parentJ] {
		(*parent)[parentI] = parentJ
	} else {
		(*parent)[parentI] = parentJ
		(*rank)[parentJ] += 1
	}
}

func find(i int, parent *[]int) int {
	if (*parent)[i] == i {
		return i
	}
	(*parent)[i] = find((*parent)[i], parent)
	return (*parent)[i]
}