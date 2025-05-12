/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return helper(root, 0)
}

func helper(node *TreeNode, sum int) int{
    if node == nil{
        return 0
    }
    sum = 10*sum + node.Val
    if node.Left == nil && node.Right == nil{
        return sum 
    }
    return helper(node.Left, sum) + helper(node.Right, sum)
    }