func rotate(nums []int, k int)  {
    n := len(nums)
    // 1, 2, 3, 4, 5, 6, 7
    // 7, 6, 5, 4, 3, 2, 1
    // 5, 6, 7, 1, 2, 3, 4
    reverseArray(nums)
    reverseArray(nums[0:k%n]) 
    reverseArray(nums[k%n: n])
}

func reverseArray(arr []int){
    i := 0
    j := len(arr)-1
    for i < j{
        temp := arr[i]
        arr[i] = arr[j]
        arr[j] = temp
        i += 1
        j -= 1 
    }
}
