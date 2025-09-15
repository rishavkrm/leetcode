func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	rank := make([]int, 26)
	dict := map[byte]int{}
	for i := range len(parent) {
		dict[byte('a'+i)] = i
		parent[i] = i
	}
	for _, eq := range equations {
		a := dict[eq[0]]
		b := dict[eq[3]]
		c := eq[1:3]
        // fmt.Println(a, b, c)
		if c == "==" {
			union(a, b, &parent, &rank)
		}
	}
	for _, eq := range equations {
		a := dict[eq[0]]
		b := dict[eq[3]]
		c := eq[1:3]
		if c == "==" {
			if find(a, &parent) != find(b, &parent) {
				return false
			}
		} else {
			if find(a, &parent) == find(b, &parent) {
				return false
			}
		}
	}
	return true
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