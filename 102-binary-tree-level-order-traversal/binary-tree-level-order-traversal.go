/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    treeQ := make([]*TreeNode, 0)
    resultQ := make([][]int, 0)
    // append root value
    treeQ = append(treeQ, root)
    // resultQ = append(resultQ, []int {*root.Val})
    return helper(treeQ, resultQ)

}

func helper(treeQ []*TreeNode, resultQ [][]int)([][]int){
    curr_len := len(treeQ)
    if curr_len == 0{
        return resultQ
    }
    curr_ResultQ := make([]int, 0)
    for i := 0; i < curr_len; i++{
        // dequeue the node
        curr_tree := treeQ[i]
        curr_ResultQ = append(curr_ResultQ, curr_tree.Val)
        // append left and right of current dequeue node
        if curr_tree.Left != nil{
            treeQ = append(treeQ, curr_tree.Left)
        }
        if curr_tree.Right != nil{
            treeQ = append(treeQ, curr_tree.Right)
        }
    }
    // dequeue all past elements
    treeQ = treeQ[curr_len:]
    // append the current to result queue
    resultQ = append(resultQ, curr_ResultQ)
    return helper(treeQ, resultQ)
}