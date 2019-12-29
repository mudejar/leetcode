package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	currNext int
	inorder  []int
}

func Constructor(root *TreeNode) BSTIterator {
	var result []int
	inorder(root, &result)
	return BSTIterator{currNext: -1, inorder: result}
}

func inorder(n *TreeNode, output *[]int) {
	if n != nil {
		inorder(n.Left, output)
		*output = append(*output, n.Val)
		inorder(n.Right, output)
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.currNext+1 < len(this.inorder) {
		this.currNext = this.currNext + 1
	}
	return this.inorder[this.currNext]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.currNext+1 < len(this.inorder)
}
