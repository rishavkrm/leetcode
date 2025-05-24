/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    return lca(root, p, q)
}

func lca(root, p, q *TreeNode) *TreeNode {
    if root == nil{
        return nil
    }
    if root.Val == p.Val || root.Val == q.Val {
        return root
    }
    left := lca(root.Left, p, q)
    right := lca(root.Right, p, q)
    if left == nil{
        return right
    }
    if right == nil{
        return left
    }
    return root

}



