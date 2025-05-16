func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)
    uniqueInts := make([]int, 0)
    for i := 0; i < len(nums); i++{
        freqMap[nums[i]] = freqMap[nums[i]] + 1
    }
    heap := make([]int, 0)
    i := 0
    for key, _ := range freqMap{
        if i < k{
            heap = append(heap, key)
        }
        uniqueInts = append(uniqueInts, key)
        i += 1
    }
    heap = MakeMinHeap(heap, freqMap)
    for i := k;  i < len(uniqueInts); i++{
        if freqMap[heap[0]] < freqMap[uniqueInts[i]]{
            heap[0] = uniqueInts[i]
            MinHeapify(heap, k, 0, freqMap)
        }
    }
    return heap
}

func MakeMinHeap(nums []int, freqMap map[int]int)[]int{
    n := len(nums)
    i := n/2-1
    for i >= 0{
        MinHeapify(nums, n, i, freqMap)
        i -= 1
    }
    return nums
}

func MinHeapify(nums []int, n, i int,freqMap map[int]int){
    smallest := i
    left := 2 * i + 1
    right := 2 * i + 2
    if left < n && freqMap[nums[smallest]] > freqMap[nums[left]]{
        smallest = left
    }
    if right < n && freqMap[nums[smallest]] > freqMap[nums[right]]{
        smallest = right
    }
    if smallest != i{
        nums[i], nums[smallest] = nums[smallest], nums[i]
        MinHeapify(nums, n, smallest, freqMap)
    }
}