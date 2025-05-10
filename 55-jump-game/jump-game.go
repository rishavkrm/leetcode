func canJump(nums []int) bool {
    curr_max := nums[0]
    n := len(nums)
    if n == 0{
        return true
    }
    for i := range(n){
        if i == 0 {
            continue
        }
        curr_max -= 1
        if curr_max < 0{
            return false
        }
        curr_max = Max(nums[i], curr_max)
    }   
    return true
}

func Max(a int, b int) int {
    if a > b{
        return a
    }
    return b
}