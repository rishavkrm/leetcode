/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{
        return nil
    }
    queue := []*TreeNode{root}
    result := make([][]int, 0)
    return helper(queue, result)
}

func helper(queue []*TreeNode, result [][]int) [][]int{
    if len(queue) == 0{
        return result
    }
    curr := make([]int, 0)
    length := len(queue)
    for i := 0; i < length; i++{
        node := queue[i]
        if node.Left != nil{
            queue = append(queue, node.Left)
        }
        if node.Right != nil{
            queue = append(queue, node.Right)
        }
        curr = append(curr, node.Val)
    }
    queue = queue[length:]
    if len(result) % 2 != 0{
        curr = reverse(curr)
    }
    result = append(result, curr)
    return helper(queue, result)
}

func reverse(arr []int) []int{
    i := 0
    j := len(arr) - 1
    for i < j{
        temp := arr[i]
        arr[i] = arr[j]
        arr[j] = temp
        i += 1
        j -= 1
    }
    return arr
}