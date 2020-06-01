package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}

	if root == nil {
		return levels
	}

	helper(root, 1, &levels)
	return levels
}

func helper(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}

	if len(*res) < level { // if this is an unallocated level of the tree
		*res = append(*res, []int{node.Val})
	} else {
		// append the current node value
		(*res)[level-1] = append((*res)[level-1], node.Val)
	}

	// process child nodes for the next level
	helper(node.Left, level+1, res)
	helper(node.Right, level+1, res)
}
