func sortColors(nums []int)  {
    low, mid, high := 0, 0, len(nums)-1
    for mid <= high{
        curr := nums[mid]
        if curr == 0{
            nums[mid], nums[low] = nums[low], nums[mid]
            mid += 1
            low += 1
        }
        if curr == 1{
            mid += 1
        }
        if curr == 2{
            nums[mid], nums[high] = nums[high], nums[mid]
            high -= 1
        }
    }

}