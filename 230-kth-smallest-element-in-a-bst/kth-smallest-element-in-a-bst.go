/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    i := 1
    return helper(root, k, &i)
}

func helper(node *TreeNode, k int, i *int) int{
    if node == nil{
        return -1
    }
    
    x := helper(node.Left, k, i)
    if *i == k{
        *i = *i+1
        return node.Val
    } 
    *i = *i+1
    y := helper(node.Right, k, i)
    fmt.Println(x, y)
    if x >= 0{
        return x
    }
    if y >= 0{
        return y
    }
    return -1

}