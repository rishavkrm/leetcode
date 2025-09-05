func longestConsecutive(nums []int) int {
	parent := make([]int, len(nums))
	rank := make([]int, len(nums))
	hash := map[int]int{}
	for i := range len(parent) {
		parent[i] = i
		_, exists := hash[nums[i]]
		if !exists {
			hash[nums[i]] = i

		}
	}
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		_, exists := hash[curr+1]
		if exists {
			union(hash[curr], hash[curr+1], &parent, &rank)
		}
		_, exists = hash[curr-1]
		if exists {
			union(hash[curr], hash[curr-1], &parent, &rank)
		}
	}
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		_, exists := hash[curr+1]
		if exists {
			union(hash[curr], hash[curr+1], &parent, &rank)
		}
		_, exists = hash[curr-1]
		if exists {
			union(hash[curr], hash[curr-1], &parent, &rank)
		}
	}
	maxi := 0
	nHash := map[int]int{}
	for _, i := range parent {
		nHash[i] += 1
		maxi = max(maxi, nHash[i])
	}
	return maxi
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