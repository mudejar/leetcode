package main

import (
	log "github.com/sirupsen/logrus"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	log.WithField("[1,1]", isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Right: nil,
			Left:  nil,
		},
		Right: nil,
	})).Println()
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	arr := []int{}
	traverseBST(root, &arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}

	return true
}

func traverseBST(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	traverseBST(node.Left, arr)
	*arr = append(*arr, node.Val)
	traverseBST(node.Right, arr)
}
