package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}

	arr = append(arr, postorderTraversal(root.Left)...)
	arr = append(arr, postorderTraversal(root.Right)...)
	arr = append(arr, root.Val)
	return arr
}
