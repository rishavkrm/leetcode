/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return maxDepthHelper(root)
}

func maxDepthHelper(node *TreeNode) int{
    if node == nil {
        return 0
    }
    x := 1 + maxDepthHelper(node.Left)
    y := 1 + maxDepthHelper(node.Right)
    return Max(x, y)
}

func Max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}