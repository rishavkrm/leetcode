func longestCommonPrefix(strs []string) string {
	m := len(strs[0])
	for i := 0; i < len(strs); i++ {
		m = min(m, len(strs[i]))
	}
    n := 0
	for i := 0; i < m; i++ {
        curr := strs[0][i]
		for j := 0; j < len(strs); j++ {
            if strs[j][i] != curr{
                return strs[0][0:n]
            }
		}
        n += 1
	}
    return strs[0][0:n]
}