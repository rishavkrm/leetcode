func minWindow(s string, t string) string {
	res := ""
	start := 0
	end := 0
	minLength := len(s) + 100

	T := make(map[byte]int)
	S := make(map[byte]int)

	count := 0

	for i := range len(t) {
		T[t[i]] += 1
	}

	for end < len(s) {
		curr := s[end]
		S[curr] += 1
		if T[curr] > 0 && S[curr] == T[curr] {
			count += 1
			if count >= len(T) {
				for start <= end {
					minLength = min(end-start+1, minLength)
					if minLength == end-start+1 {
						res = s[start : end+1]
					}
					S[s[start]] -= 1
					if S[s[start]] < T[s[start]] {
						start += 1
						count -= 1
						break
					}
					start += 1
				}
			}
		}
		end += 1

	}
	return res
}