func subarraySum(nums []int, k int) int {
	p := make([]int, len(nums)+1)
	curr := 0
	for i := range len(p) {
		p[i] = curr
        if i == len(p)-1{
            break
        }
		curr += nums[i]

	}
    c := 0
	for i := 0; i < len(p); i++{
        for j := i+1; j < len(p); j++{
            if p[j]-p[i] == k{
                c += 1
            }
        }
    }
    return c
}
