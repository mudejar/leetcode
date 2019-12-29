package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	arr := []*TreeNode{}
//	traverseBST(root, &arr)
//
//	for _, node := range arr {
//		if node.Val > p.Val {
//			return node
//		}
//	}
//
//	return nil
//}
//
//func traverseBST(node *TreeNode, arr *[]*TreeNode) {
//	if node == nil {
//		return
//	}
//
//	traverseBST(node.Left, arr)
//	*arr = append(*arr, node)
//	traverseBST(node.Right, arr)
//}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var closest *TreeNode // maintain the closest node
	cur := root
	for cur != nil {
		if cur.Val > p.Val {
			closest = cur // since cur node is higher than p, it could be potential result as we move closer
			cur = cur.Left
		} else if cur.Val <= p.Val { // if we have found the value, then result is either in closest or in right subtree
			cur = cur.Right
		}
	}
	return closest
}
