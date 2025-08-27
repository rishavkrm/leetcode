func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	maxWater := 0
	for left < right {
		maxWater = max(maxWater, min(height[left], height[right])*(right-left))
        if height[left] < height[right]{
            left += 1
        }else{
            right -= 1
        }
	}
    return maxWater
}