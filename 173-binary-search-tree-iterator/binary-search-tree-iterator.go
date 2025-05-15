/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    stack := make([]*TreeNode, 0)
    curr := root
    for curr != nil{
        stack = append(stack, curr)
        curr = curr.Left
    }
    return BSTIterator{
        stack : stack,
    }

}


func (this *BSTIterator) Next() int {
    stack := this.stack
    popped := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    curr := popped.Right
    for curr != nil{
        stack = append(stack, curr)
        curr = curr.Left
    }
    this.stack = stack
    return popped.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

 // []