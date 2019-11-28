package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var count int

func postOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	s1, s2 := postOrder(root.Left), postOrder(root.Right)
	if s1 == true && s2 == true {
		if root.Left != nil && root.Left.Val != root.Val {
			return false
		}
		if root.Right != nil && root.Right.Val != root.Val {
			return false
		}
		count++
		return true
	}
	return false
}

func countUnivalSubtrees(root *TreeNode) int {
	count = 0
	postOrder(root)
	return count
}
