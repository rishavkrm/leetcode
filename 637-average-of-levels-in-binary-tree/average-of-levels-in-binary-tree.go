/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil{
        return nil
    }
    queue := []*TreeNode{root}
    result := make([]float64, 0)
    return helper(queue, result)
}

func helper(queue []*TreeNode, result []float64) []float64{
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
    avg := average(curr)
    result = append(result, avg)
    return helper(queue, result)
}

func average(arr []int) float64{
    sum := 0
    for i := 0; i < len(arr); i++{
        sum += arr[i]
    }
    return float64(sum)/float64(len(arr))
} 