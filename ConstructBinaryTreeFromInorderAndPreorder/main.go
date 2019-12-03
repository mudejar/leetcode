package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	val := preorder[0]
	leftTreeSize := 0
	for i, v := range inorder {
		if v == val {
			leftTreeSize = i
		}
	}
	return &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:leftTreeSize+1], inorder[:leftTreeSize]),
		Right: buildTree(preorder[leftTreeSize+1:], inorder[leftTreeSize+1:]),
	}
}
