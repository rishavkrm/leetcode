/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := Length(head)
	p := n / k
	curr := head
	res := head
	prevHead := &ListNode{}
	for j := range p { // p = 0, j = 1
		currHead := curr        // curr iteration head
		nextHead := &ListNode{} // next iteration head
		currTail := &ListNode{} // curr iteration tail
		prev := &ListNode{}
		for i := 0; i < k; i++ { // 0
			next := curr.Next // 4 5
			nextHead = next   // 4 5
			curr.Next = prev  // 4->3->nil
			prev = curr       // 3 4
			currTail = curr          // 3 4
			curr = next       // 4 5

		}
		if j == 0 { // make tail of 1st iteration as head
			res = prev
		}
		currHead.Next = nextHead // curr iteration head bacame tail now, its next = next iteration head
		if prevHead != nil {
			prevHead.Next = currTail 
		}
		prevHead = currHead
	}
	return res
}

func Length(head *ListNode) int {
	count := 0
	curr := head
	for curr != nil {
		count += 1
		curr = curr.Next
	}
	return count
}