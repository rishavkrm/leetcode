/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, l int, r int) *ListNode {
	var c1 *ListNode
	var c2 *ListNode
	var prev *ListNode
	curr := head
	xx := 1
	yy := 1
	for curr != nil && (c1 == nil || c2 == nil) {
		if xx == l {
			c1 = prev
		}
		if yy == r {
			c2 = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
		xx += 1
		yy += 1
	}
	if c1 == nil {
		temp2 := head
		head = curr
		curr = temp2
	} else {
		curr = c1.Next
	}
	prev = c2
	for curr != c2 {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	if c1 != nil {
		c1.Next = prev
	}
	return head
}