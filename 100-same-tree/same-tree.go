/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return EqualityHelper(p, q)
}

func EqualityHelper(p *TreeNode, q *TreeNode) bool{

    // check if p/q is/are nil are decide accordingly
    if p == nil && q == nil{
        return true
    }else if p == nil && q != nil{
        return false
    }else if p != nil && q == nil{
            return false
    }
    // check value of p and q
    if p.Val == q.Val {
        return EqualityHelper(p.Left, q.Left) && EqualityHelper(p.Right, q.Right)
        
    }else{
        return false
    }
}