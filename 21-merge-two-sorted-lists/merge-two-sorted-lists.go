/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				curr.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			} else {
				curr.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			}
		} else {
			if l1 != nil {
				curr.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				curr.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
		}
        curr = curr.Next
	}
	return res.Next

}