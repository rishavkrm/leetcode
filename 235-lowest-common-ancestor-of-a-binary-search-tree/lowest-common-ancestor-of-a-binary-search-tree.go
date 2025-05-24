/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 0, 2, 3, 4, 5, 6, 7, 8, 9
    return lcaBST(root, p, q) 
}

func lcaBST(node, p, q *TreeNode) *TreeNode {
    if node == nil{
        return node
    }
    if node.Val == p.Val || node.Val == q.Val{
        return node
    }

    if node.Val > p.Val && node.Val > q.Val{
        return lcaBST(node.Left, p, q)
    }
    if node.Val < p.Val && node.Val < q.Val{
        return lcaBST(node.Right, p, q)
    }
    return node
}