/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	helper(root, targetSum, []int{}, &result)
	return result
}
func helper(curr *TreeNode, targetSum int, arr []int, result *[][]int) {
	if curr == nil {
		return
	}
	arr = append(arr, curr.Val)
	targetSum -= curr.Val
	fmt.Println(arr, 22-targetSum)

	if targetSum == 0 && curr.Left == nil && curr.Right == nil {
		*result = append(*result, copyArray(arr))
	}

	helper(curr.Left, targetSum, arr, result)
	helper(curr.Right, targetSum, arr, result)
}

func copyArray(arr []int)[]int{
    res := make([]int, len(arr))
    for i := range len(res){
        res[i] = arr[i]
    }
    return res
}