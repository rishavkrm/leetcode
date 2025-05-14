func findKthLargest(nums []int, k int) int {
    // first make a heap of size k
    n := len(nums)
    heap := nums[:k]
    heap = MakeHeap(heap)
    if n <= k{
        return heap[0]
    }
    for i := k; i < n; i++{
        if heap[0] >= nums[i]{
            continue
        }
        heap[0] = nums[i]
        heap = Heapify(heap, k, 0)
    }
    return heap[0]
}

func Heapify(nums []int, n, i int)[]int{
    smallest := i
    left := 2 * i + 1
    right := 2 * i + 2

    if left < n && nums[left] < nums[smallest]{
        smallest = left
    }
    if right < n && nums[right] < nums[smallest]{
        smallest = right
    }
    if i != smallest{
        nums[i], nums[smallest] = nums[smallest], nums[i]
        Heapify(nums, n, smallest)
    }
    return nums
}

func MakeHeap(arr []int) []int{
    // first make it heap 
    n := len(arr)
    for i := n/2 ; i >= 0; i--{
        arr = Heapify(arr, n, i)
    }
    return arr
}