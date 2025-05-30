// 1. make a min-heap with first k elements which will store k largest elements
// 2. root of the heap is the smallest of the heap
// 3. if a number is greater than root that means it comes in heap and 
// replace it with the number and heapify down
// 4. if a number is smaller then not eligible
// 5. root will be kth largeest number

func findKthLargest(nums []int, k int) int {
    heapLength := k
    heap := nums[:heapLength]
    for i := k/2 - 1; i >= 0; i--{
        minHeapify(heap, k, i)
    }
    for i := k; i < len(nums); i++{
        if nums[i] > heap[0]{
            heap[0] = nums[i]
            minHeapify(heap, k, 0)
        }
    }
    return heap[0]
}

// make an array a min-heap
func minHeapify(nums []int, n int, i int){
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
        minHeapify(nums, n, smallest)
    }
}