func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefixProduct := make([]int, n)
    suffixProduct := make([]int, n)
    p := 1
    s := 1
    for i := 0; i < len(nums); i++{
        p = p * nums[i]
        s = s * nums[n-i-1]
        prefixProduct[i] = p
        suffixProduct[n-i-1] = s
    }
    for i := 0; i < len(nums); i++{
        if i == 0{
            nums[i] = suffixProduct[i+1]
        }else if i == n-1{
            nums[i] = prefixProduct[i-1]
        }else{
            nums[i] = prefixProduct[i-1] * suffixProduct[i+1]
        }
    }
    return nums
}