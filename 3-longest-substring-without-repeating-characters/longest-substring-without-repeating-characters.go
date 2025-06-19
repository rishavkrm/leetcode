func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[byte]int)
	start := 0
	end := 0
    minLength := 0  

	for end < len(s) {
        count, ok := hashMap[s[end]]
        if !ok || count == 0{
            hashMap[s[end]] = 1
            end += 1
            minLength = max(minLength, end - start)
        }else{
            hashMap[s[start]] -= 1
            start += 1
        }
	}
    return minLength
}