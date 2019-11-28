package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) &&
		isMirror(left.Right, right.Left) &&
		isMirror(left.Left, right.Right)
}
