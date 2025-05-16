/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    arr := make([]int, 0)
    stack := make([]*TreeNode, 0)
    curr := root

    for curr != nil || len(stack) > 0 {
        for curr != nil{
            stack = append(stack, curr)
            curr = curr.Left
        }
        popped := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        arr = append(arr, popped.Val)
        curr = popped.Right
    }
    return isStrictlyIncreasing(arr)
}

func isStrictlyIncreasing(arr []int) bool{
    for i := 0; i < len(arr)-1; i++{
        if arr[i+1] <= arr[i]{
            return false
        }
    }
    return true
}