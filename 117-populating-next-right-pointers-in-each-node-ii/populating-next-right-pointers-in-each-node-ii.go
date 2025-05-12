/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil{
        return root
    }
    q := []*Node{root}
    helper(q)
    return root

}

func helper(q []*Node){
    q_len := len(q)
    if q_len == 0{
        return 
    }
    for i := 0; i < q_len; i++{
        deq := q[i]
        if i < q_len-1{
            deq.Next = q[i+1]
        }else{
            deq.Next = nil
        }
        if deq.Left != nil{
            q = append(q, deq.Left)
        }
        if deq.Right != nil{
            q = append(q, deq.Right)
        }
    }
    helper(q[q_len:])
    return
}