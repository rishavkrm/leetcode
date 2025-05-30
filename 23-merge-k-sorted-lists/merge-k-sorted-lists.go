/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1. The first element of final result will be 0th index element of any one of arrays
// 2. We will make a priority queue with the fist node of each array and then heapify it
// 3. While length of heap > 0, make the listNode as ne

func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]*ListNode, 0)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap = append(heap, lists[i])
		}
	}
	for i := len(heap)/2 - 1; i >= 0; i-- {
		minHeapify(heap, len(heap), i)
	}
    if len(heap) == 0{
        return nil
    }
	head := heap[0]
	currSortedList := head
    heapLength := len(heap) 
	for heapLength > 0 {
		if heap[0].Next != nil {
			heap[0] = heap[0].Next
		} else {
			heap[0], heap[heapLength-1] = heap[heapLength-1], heap[0]
            heapLength -= 1
		}
        if heapLength > 0{
            minHeapify(heap, heapLength, 0)
            currSortedList.Next = heap[0]
            currSortedList = currSortedList.Next
        }
		
	}
	return head
}

// make an array a min-heap
func minHeapify(nums []*ListNode, n int, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && nums[left].Val < nums[smallest].Val {
		smallest = left
	}
	if right < n && nums[right].Val < nums[smallest].Val {
		smallest = right
	}
	if i != smallest {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		minHeapify(nums, n, smallest)
	}
}
