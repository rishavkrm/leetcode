/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    return helper(root, 999999)
}

func helper(node *TreeNode, min int)int{
    if node == nil{
        return min
    }
    left := node.Left
    right := node.Right
    max_left := 999999
    min_right := 999999
    for left != nil{
        max_left = left.Val
        left = left.Right
    }
    for right != nil{
        min_right = right.Val
        right = right.Left
    }
    min = Min(min, Min(node.Val-max_left, node.Val-min_right))
    min = Min(min, helper(node.Left, min))
    min = Min(min, helper(node.Right ,min))
    return min
}

func Min(a , b int) int{
    if a < 0{
        a = -1*a
    }
    if b < 0{
        b = -1*b
    }
    if a > b{
        return b
    }
    return a
}