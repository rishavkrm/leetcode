/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil{
        return nil
    }
    curr := head
    h := &Node{}
    n := h
    newList := make([]*Node, 0)
    addressIndexMap := make(map[*Node]int)
    var prev *Node
    for curr != nil{
        n.Val = curr.Val
        n.Next = &Node{}
        addressIndexMap[curr] = len(newList)
        newList = append(newList, n)
        prev = n
        curr = curr.Next
        n = n.Next
    }   
    if prev != nil{
        prev.Next = nil
    }
    n = h
    curr = head
    for curr != nil{
        if curr.Random != nil{
            n.Random = newList[addressIndexMap[curr.Random]]
        }
        curr = curr.Next
        n = n.Next
    }
    return h


}