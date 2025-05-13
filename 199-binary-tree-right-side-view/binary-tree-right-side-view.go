/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil{
        return nil
    }
    queue := []*TreeNode{root}
    result := make([]int, 0)
    return helper(queue, result)
}

func helper(queue []*TreeNode , result []int) []int{
    if len(queue) == 0{
        return result
    }
    length := len(queue)
    for i := 0; i < length; i++{
        dequeued := queue[i]
        if dequeued.Left != nil{
            queue = append(queue, dequeued.Left)
        }
        if dequeued.Right != nil{
            queue = append(queue, dequeued.Right)
        }
        if i == length-1{
            result = append(result, dequeued.Val)
        }
    }
    queue = queue[length:]
    return helper(queue, result)
}