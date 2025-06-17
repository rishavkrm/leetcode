func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[0]))
		return res
	}
	a := 0
	b := 1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			res = addToRes(nums[a], nums[b-1], res)
			break
		}
		if nums[i]+1 == nums[b] {
			b += 1
		} else {
			res = addToRes(nums[a], nums[b-1], res)
			a = i + 1
			b += 1
		}
	}
	return res
}

func addToRes(a, b int, res []string) []string {
	if a == b {
		res = append(res, strconv.Itoa(a))
	} else {
		res = append(res, strconv.Itoa(a)+"->"+strconv.Itoa(b))
	}
	return res

}