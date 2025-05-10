func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefixSum := make([]int, n)
    suffixSum := make([]int, n)
    p := 1
    s := 1
    for i := 0; i < len(nums); i++{
        p = p * nums[i]
        s = s * nums[n-i-1]
        prefixSum[i] = p
        suffixSum[n-i-1] = s
    }
    for i := 0; i < len(nums); i++{
        if i == 0{
            nums[i] = suffixSum[i+1]
        }else if i == n-1{
            nums[i] = prefixSum[i-1]
        }else{
            nums[i] = prefixSum[i-1] * suffixSum[i+1]
        }
    }
    return nums

    
}