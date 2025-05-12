/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    invertHelper(root)
    return root
}

func invertHelper(node *TreeNode){
    if node == nil{
        return
    }
    temp := node.Right
    node.Right = node.Left
    node.Left = temp
    invertHelper(node.Left)
    invertHelper(node.Right)
}