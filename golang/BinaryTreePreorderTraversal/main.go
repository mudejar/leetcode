package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}

	arr = append(arr, root.Val)
	arr = append(arr, preorderTraversal(root.Left)...)
	arr = append(arr, preorderTraversal(root.Right)...)
	return arr
}
