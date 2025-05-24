/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int{
    maxSum := -99999
    helper(root, &maxSum)
    return maxSum
}

func helper(node *TreeNode, maxSum *int) int{
    if node == nil {
        return -99999
    }
    leftGain := max(0, helper(node.Left, maxSum))
    rightGain := max(0, helper(node.Right, maxSum))
    maxPathSum := node.Val + leftGain + rightGain
    *maxSum = max(*maxSum, maxPathSum)
    return node.Val + max(leftGain, rightGain)
}

