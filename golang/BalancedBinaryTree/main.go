package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := height(root)
	return isBalanced
}

func height(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	left, leftIsBalanced := height(node.Left)
	right, rightIsBalanced := height(node.Right)

	if !leftIsBalanced || !rightIsBalanced {
		return max(left, right) + 1, false
	}

	if abs(left, right) > 1 {
		return max(left, right) + 1, false
	}

	return max(left, right) + 1, true
}

func abs(x, y int) int {
	z := x - y
	if z < 0 {
		return -z
	}
	return z
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
