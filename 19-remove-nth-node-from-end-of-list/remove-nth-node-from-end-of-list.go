/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    curr := head
    hashmap := make(map[int]*ListNode);
    len := 1;
    for curr != nil{
        hashmap[len] = curr
        curr = curr.Next
        len += 1
    }    
    if n == len-1{
        return head.Next
    }
    hashmap[len-n-1].Next = hashmap[len-n+1]
    return head
    
}