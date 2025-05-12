/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return symmetricHelper(root.Left, root.Right)
}

func symmetricHelper(nodeA *TreeNode, nodeB *TreeNode) bool{
    if nodeA == nil && nodeB == nil{
        return true
    }
    if nodeA == nil && nodeB != nil || nodeA != nil && nodeB == nil{
        return false
    }
    if nodeA.Val == nodeB.Val{
        return symmetricHelper(nodeA.Left, nodeB.Right) && symmetricHelper(nodeA.Right, nodeB.Left) 
    } else{
        return false
    }
}