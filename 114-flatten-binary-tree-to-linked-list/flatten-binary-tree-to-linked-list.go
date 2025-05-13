/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    helper(root)
}

func helper(node *TreeNode) {
    if node == nil{
        return
    }
    left := node.Left
    for left != nil && left.Right != nil{
        if left.Right != nil{
            left = left.Right
        }
    }
    temp := node.Right
    if node.Left != nil{
        node.Right = node.Left
        left.Right = temp
        node.Left = nil
    }
    helper(node.Right)
}