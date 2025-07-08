func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	res := [][]string{}
	for i := range len(strs) {
		curr := strs[i]
		dict[sortedString(curr)] = append(dict[sortedString(curr)], curr)
	}
    for _, val := range(dict){
        res = append(res, val)
    }
    return res
}

func sortedString(input string) string {
	bytes := []byte(input)
	sort.Slice(bytes, func(a, b int) bool {
		return bytes[a] < bytes[b]
	})
	return string(bytes)
}