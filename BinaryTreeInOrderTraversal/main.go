package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}

	arr = append(arr, inorderTraversal(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)
	return arr
}
