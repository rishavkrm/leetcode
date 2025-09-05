func accountsMerge(accounts [][]string) [][]string {
	emails := map[string]int{}
	count := 0
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			email := account[i]
			_, exists := emails[email]
			if !exists {
				emails[email] = count
				count += 1
			}
		}
	}
	name := make([]int, len(emails))
	rank := make([]int, len(emails))
	parent := make([]int, len(accounts))
	for i, _ := range parent {
		parent[i] = i
	}
	for x, account := range accounts {
		for i := 1; i < len(account); i++ {
			email := account[i]
			name[emails[email]] = x
		}
	}
	for x, account := range accounts {
		for i := 1; i < len(account); i++ {
			nme := name[emails[account[i]]]
			if nme != x { // 4->1
				union(x, nme, &parent, &rank)
			}
		}
	}
    for x, account := range accounts {
		for i := 1; i < len(account); i++ {
			nme := name[emails[account[i]]]
			if nme != x { // 4->1
				union(x, nme, &parent, &rank)
			}
		}
	}
	dict := map[int][]string{}
	result := [][]string{}
	for i := range parent {
		dict[parent[i]] = append(dict[parent[i]], (accounts[i])[1:]...)
	}
	dict2 := map[string]int{}
	for key, val := range dict {
		temp := []string{}
		for _, v := range val {
			if dict2[v] == 0 {
				temp = append(temp, v)
				dict2[v] = 1
			}
		}
        sort.Slice(temp, func(i int, j int) bool { return temp[i] < temp[j] })
        temp = append([]string{accounts[key][0]}, temp...)
		result = append(result, temp)
	}
	return result

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
