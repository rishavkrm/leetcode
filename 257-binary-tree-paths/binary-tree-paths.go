/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	helper(root, "", &result)
	return result
}

func helper(root *TreeNode, curr string, result *[]string) {
	if root == nil {
		return
	}
	if len(curr) > 0 {
		curr += "->"
	}
	curr += fmt.Sprintf("%d", root.Val)
	if root.Left == nil && root.Right == nil {
		(*result) = append((*result), curr)
		return
	}
	helper(root.Left, curr, result)
	helper(root.Right, curr, result)
}