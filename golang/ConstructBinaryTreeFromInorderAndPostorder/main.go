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
	inOrder := []int{9, 3, 15, 20, 7}
	postOrder := []int{9, 15, 7, 20, 3}
	log.WithFields(
		log.Fields{
			"In order":           inOrder,
			"Post order":         postOrder,
			"buildTree response": buildTree(inOrder, postOrder),
		},
	).Println()
}

var (
	inPos   map[int]int
	inorder []int
	post    []int
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	inPos = make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	post = postorder
	return buildTreeHelper(0, len(postorder)-1, 0)
}

func buildTreeHelper(postStart int, postEnd int, inStart int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	root := &TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildTreeHelper(postStart, postStart+leftLen-1, inStart)
	root.Right = buildTreeHelper(postStart+leftLen, postEnd-1, rootIdx+1)

	return root
}
