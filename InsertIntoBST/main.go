package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	if node == nil {
		return &TreeNode{
			Val: val,
		}
	}

	for node != nil {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{
					Val: val,
				}
				return root
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{
					Val: val,
				}
				return root
			}
			node = node.Right
		}
	}

	return node
}