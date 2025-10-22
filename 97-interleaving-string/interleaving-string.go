func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]int, len(s1))
	for i := range len(dp) {
		dp[i] = make([]int, len(s2))
	}
	return helper(s1, s2, s3, 0, 0, dp)
}

func helper(s1 string, s2 string, s3 string, i1 int, i2 int, dp [][]int) bool {
	// if i1 == len(s1)-1 && i2 == len(s2)-1 {
	// 	return true
	// }
	i3 := i1 + i2
	if i1 == len(s1) || i2 == len(s2) {
		if i1 == len(s1) && i2 == len(s2) {
            fmt.Println("Yes")
			return true
		}
		if i1 == len(s1) {
			if s2[i2:] == s3[i3:] {
				return true
			} else {
				return false
			}
		} else {
			if s1[i1:] == s3[i3:] {
				return true
			} else {
				return false
			}
		}
	}
	if dp[i1][i2] != 0 {
		return dp[i1][i2] == 1
	}

	x, y := false, false
	if s1[i1] == s3[i3] && s2[i2] == s3[i3] {
		x = helper(s1, s2, s3, i1+1, i2, dp)
		y = helper(s1, s2, s3, i1, i2+1, dp)
	} else if s1[i1] == s3[i3] {
		x = helper(s1, s2, s3, i1+1, i2, dp)
	} else if s2[i2] == s3[i3] {
		y = helper(s1, s2, s3, i1, i2+1, dp)
	}
	if x || y {
		dp[i1][i2] = 1
	} else {
		dp[i1][i2] = -1
	}
	return x || y
}