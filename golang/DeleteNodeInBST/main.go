package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func deleteNode(root *TreeNode, key int) *TreeNode {
	cur, par := findNode(root, nil, key)
	if cur == nil {
		// the node is already deleted
		return root
	}

	if par == nil {
		// this happens when the node to delete is root itself
		return adjustTree(cur)
	}

	if par.Left == cur {
		par.Left = adjustTree(cur)
	} else {
		par.Right = adjustTree(cur)
	}

	return root
}

// findNode returns a tuple of the found node and its parent
func findNode(t, par *TreeNode, key int) (*TreeNode, *TreeNode) {
	if t == nil {
		return nil, par
	}

	if t.Val == key {
		return t, par
	}

	if t.Val < key {
		return findNode(t.Right, t, key)
	}

	return findNode(t.Left, t, key)
}

// adjustTree deletes the given node and adjusts the subtrees under the given node.
func adjustTree(t *TreeNode) *TreeNode {
	if t.Left == nil {
		return t.Right
	}

	if t.Right == nil {
		return t.Left
	}

	// This is the hard part, you need to find the correct node to replace the
	// head with. At his point we know that the neither the left nor right
	// children are nil
	root := t.Left // you have to assign this because you need to return root
	cur := root

	// Go to the furthest right child under the original left child we looked at
	for cur.Right != nil {
		cur = cur.Right
	}

	// After we reach the right most node, assign the right child of this node
	// to be the right child of the root node of this subtree
	cur.Right = t.Right

	return root
}
