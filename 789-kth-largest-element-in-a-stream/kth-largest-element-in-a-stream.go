type KthLargest struct {
	heap []int
    k int
}

func Constructor(k int, nums []int) KthLargest {
    for len(nums) < k{
        nums = append(nums, -9999)
    }
    n := len(nums)
    nums = HeapSort(nums)
    return KthLargest{
        heap : nums[n-k:],
        k : k,
    }
}
func Heapify(arr []int, n, i int) []int {
    largest := i
	left := 2 * i + 1
	right := 2 * i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i{
        arr[largest], arr[i] = arr[i], arr[largest]
        Heapify(arr, n, largest)
    }
    return arr
}

func HeapSort(arr []int) []int{
    n := len(arr)
    // build heap
    for i := n/2-1; i >= 0; i--{
		arr = Heapify(arr, n, i)
	}
    for i := n-1; i > 0; i--{
        temp := arr[0]
        arr[0] = arr[i]
        arr[i] = temp
        arr = Heapify(arr, i, 0)
    } 
    return arr
}

func (this *KthLargest) Add(val int) int {
    heap := this.heap
    k := this.k
    // [1, 2, 3, 5]
    for i := 0; i<k; i++{
        if val >= heap[i]{
            if i > 0{
                heap[i], heap[i-1] = heap[i-1], heap[i]
            }else{
                heap[i] = val
            }
        } else{
            break
        }
    }
    this.heap = heap
    return heap[0]

    
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */