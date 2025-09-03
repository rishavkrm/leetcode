/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		if curr.Val == val {
			if prev != nil {
				prev.Next = curr.Next
				curr = curr.Next
			} else {
				head = head.Next
				curr = head
			}
			continue
		}
		prev = curr
		curr = curr.Next
	}
	return head
}