/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	dp1 := map[*TreeNode]int{}
	dp2 := map[*TreeNode]int{}

	return helper(root, false, dp1, dp2)
}

func helper(root *TreeNode, parentTaken bool, dp1 map[*TreeNode]int, dp2 map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if !parentTaken {
		val, exists := dp1[root]
		if exists {
			return val
		}
		case1 := root.Val + helper(root.Left, true, dp1, dp2) + helper(root.Right, true, dp1, dp2)
		case2 := helper(root.Left, false, dp1, dp2) + helper(root.Right, false, dp1, dp2)
		dp1[root] = max(case1, case2)
		return dp1[root]
	}
	case2 := helper(root.Left, false, dp1, dp2) + helper(root.Right, false, dp1, dp2)
	val, exists := dp2[root]
	if exists {
		return val
	}
    dp2[root] = case2
	return case2
}