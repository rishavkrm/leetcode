func sortColors(nums []int) {
	red := 0
	blue := len(nums) - 1
	for i, num := range nums {
		if num == 0 {
			if nums[red] == 0 {
				red += 1
				continue
			}
			temp := nums[red]
			nums[red] = 0
			nums[i] = temp
			red += 1
		}
	}
	for i, _ := range nums {
		num := nums[len(nums)-1-i]
		if num == 2 {
			if nums[blue] == 2 {
				blue -= 1
				continue
			}
			temp := nums[blue]
			nums[blue] = 2
			nums[len(nums)-i-1] = temp
			blue -= 1
		}
	}
}