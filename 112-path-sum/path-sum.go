/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }
    return helper(root,0, targetSum)
}

func helper(node *TreeNode, sum int, targetSum int) bool{
    if node == nil{
        return false
    }
    sum += node.Val
    if node.Left == nil && node.Right == nil{
        return sum == targetSum
    } 
    return helper(node.Left, sum, targetSum) || helper(node.Right, sum, targetSum)
}