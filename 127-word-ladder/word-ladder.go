func ladderLength(beginWord string, endWord string, wordList []string) int {
	hash := map[string]bool{}
	for _, word := range wordList {
		hash[word] = true
	}
	if !hash[beginWord] {
		wordList = append(wordList, beginWord)
		hash[beginWord] = true
	}
	start := -1
	end := -1
	for i := range len(wordList) {
		if wordList[i] == endWord {
			end = i
		}
		if wordList[i] == beginWord {
			start = i
		}
	}
	if end == -1 || start == -1 {
		return 0
	}
	graph := map[int][]int{}
    for i := range len(hash){
        graph[i] = []int{}
    }
	for i := range len(hash) {
		for j := range len(hash) {
			if i == j {
				continue
			}
			if isDifferenceOne(wordList[i], wordList[j]) {
				graph[i] = append(graph[i], j)
			}
		}
	}
	x :=  bfs(graph, start, end) 
	if x == 10000000{
        return 0
    }
    return x+1
}

func bfs(graph map[int][]int, start int, end int) int {
	v := map[int]bool{}
	dist := map[int]int{}
	for i := range len(graph) {
		dist[i] = 10000000
	}
	dist[start] = 0
	curr := start
	for len(v) != len(graph) {
		v[curr] = true
		for _, nbr := range graph[curr] {
			if !v[nbr] {
				dist[nbr] = min(dist[nbr], dist[curr]+1)
			}
		}
		min := 10000000
		next := -1
		for i,_ := range graph {
			if !v[i] {
                if dist[i] < min{
                    min = dist[i]
                    next = i
                }
			}
		}
        if next == -1{
            break
        }
        curr = next
	}
    return dist[end]

}

func isDifferenceOne(s1 string, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff += 1
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}