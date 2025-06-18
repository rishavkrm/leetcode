/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	curr := head 
	prev := &ListNode{} 
	res := head.Next
	for curr != nil && curr.Next != nil {
		next := curr.Next
		if prev != nil {
			prev.Next = next 
		}
		curr.Next = next.Next 
		next.Next = curr 
		prev = curr 
		curr = curr.Next
	}
	return res
}