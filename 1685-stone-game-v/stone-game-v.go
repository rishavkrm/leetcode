func stoneGameV(stoneValue []int) int {
	sumArr := make([]int, len(stoneValue))
	curr := 0
    dp := make([][]int, len(stoneValue))
    for i := 0; i < len(dp); i++{
        dp[i] = make([]int, len(dp))
    }
	for i := 0; i < len(sumArr); i++ {
		curr += stoneValue[i]
		sumArr[i] = curr
	}
    for i := 0; i < len(dp); i++{
        for j := 0; j < len(dp[0]); j++{
            dp[i][j] = -1
        }
    }
	return helper(stoneValue, 0, len(stoneValue)-1, sumArr, dp)
}

func helper(arr []int, left int, right int, sumArr []int, dp [][]int) int {
	if left >= right {
		return 0
	}
	maxi := 0
    if dp[left][right] != -1{
        return dp[left][right]
    }
 	for i := left; i < right; i++ {
		res := 0
		sumL := sumArr[i]
		if left > 0 {
			sumL = sumArr[i] - sumArr[left-1]
		}
		sumR := sumArr[right] - sumArr[i]
		if sumL > sumR {
			res = sumR + helper(arr, i+1, right, sumArr, dp)
		} else if sumL < sumR {
			res = sumL + helper(arr, left, i, sumArr, dp)
		} else {
			res1 := sumR + helper(arr, i+1, right, sumArr, dp)
			res2 := sumL + helper(arr, left, i, sumArr, dp)
			res = max(res1, res2)
		}
		maxi = max(maxi, res)
	}
    dp[left][right] = maxi
	return maxi
}

// [6, 8, 11, 15, 20, 25]

// 0, 5