/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    arr []int
    pointer int
}


func Constructor(root *TreeNode) BSTIterator {
    arr := make([]int, 0)
    arr = inorder(root, arr)
    fmt.Println(arr)
    return BSTIterator{
        arr : arr,
        pointer: 0,
    }

}


func (this *BSTIterator) Next() int {
    x := this.arr[this.pointer]
    this.pointer += 1
    return x
}


func (this *BSTIterator) HasNext() bool {
    if this.pointer < len(this.arr){
        return true
    } 
    return false
}

func inorder(node *TreeNode, arr []int)[]int{
    if node == nil{
        return arr
    }
    arr = inorder(node.Left, arr)
    arr = append(arr, node.Val)
    arr = inorder(node.Right, arr)
    return arr
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */