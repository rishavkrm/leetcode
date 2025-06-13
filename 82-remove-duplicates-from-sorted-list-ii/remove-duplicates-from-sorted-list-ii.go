/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	currNode := head
	prev := &ListNode{}
	for currNode != nil {
		if currNode.Next != nil && currNode.Val == currNode.Next.Val {
			temp := currNode.Val
			for currNode != nil && currNode.Val == temp {
				currNode = currNode.Next
			}
            if temp == head.Val{
                head = currNode
            }
            prev.Next = currNode;
            continue
		}
		prev = currNode          // 2
		currNode = currNode.Next // 3
	}
	return head
}