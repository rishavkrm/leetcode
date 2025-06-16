func strStr(haystack string, needle string) int {
	i := 0
	j := len(needle)

	for j <= len(haystack) {
        if haystack[i:j] == needle{
            return i
        }
        i += 1
        j += 1
	}
    return -1
}