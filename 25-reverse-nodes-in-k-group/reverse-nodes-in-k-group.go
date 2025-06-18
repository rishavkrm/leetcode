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
		x := curr //3
		y := &ListNode{}
		z := &ListNode{}
		prev := &ListNode{}
		for i := 0; i < k; i++ { // 0
			next := curr.Next // 4 5
			y = next          // 4 5
			curr.Next = prev  // 4->3->nil
			prev = curr       // 3 4
			z = curr          // 3 4
			curr = next       // 4 5

		}
		if j == 0 {
			res = prev
		}
		x.Next = y // 4->3->5
		if prevHead != nil {
			prevHead.Next = z
		}
        prevHead = x
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