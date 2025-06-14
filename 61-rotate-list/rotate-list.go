/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 0
	curr := head
	for curr != nil {
		if curr.Next == nil {
			curr.Next = head
            n += 1
			break
		}
        n += 1
        curr = curr.Next

	}
	k = k % n
	count := 0
    tail := n - k - 1
	curr = head
	for tail >= 0{
		if count == tail {
			res := curr.Next
			curr.Next = nil
			return res
		}
		count += 1
		curr = curr.Next
	}
	return curr
}