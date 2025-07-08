func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	res := [][]string{}
	for i := range len(strs) {
		curr := strs[i]
		dict[makeIdOfString(curr)] = append(dict[makeIdOfString(curr)], curr)
	}
	for _, val := range dict {
		res = append(res, val)
	}
	return res
}

func makeIdOfString(input string) string {
	dict := make([]byte, 26)
	for i := 0; i < len(input); i++ {
		dict[input[i]-97] += 1
	}
	return string(dict)
}