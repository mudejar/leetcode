package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans = 1
	depth(root)

	return ans - 1
}


func depth(t *TreeNode) int {
	if t == nil {
		return 0
	}

	L := depth(t.Left)
	R := depth(t.Right)
	ans = max(ans, L+R+1)
	return max(L,R) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}