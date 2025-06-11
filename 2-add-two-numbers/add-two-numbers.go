/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	carry := 0
	curr := head
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
		}
        sum := (n1 + n2 + carry)%10
        carry = (n1 + n2 + carry)/10
        curr.Next = &ListNode{Val : sum,}
        curr = curr.Next
        if l1 != nil{
            l1 = l1.Next
        }
        if l2 != nil{
            l2 = l2.Next
        }
	}
    if carry > 0{
        curr.Next = &ListNode{
            Val : carry,
        }
    }
    return head.Next
}
