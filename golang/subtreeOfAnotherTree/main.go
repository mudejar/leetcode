package main

import (
	"fmt"
	"strings"
)

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isSubtree(s *TreeNode, t *TreeNode) bool {
//	tree := preorder(s)
//	subtree := preorder(t)
//	return strings.Contains(tree, subtree)
//}
//
//func preorder(t *TreeNode) string {
//	if t == nil {
//		return "nil"
//	}
//	return fmt.Sprintf("#%d", t.Val) + " " + preorder(t.Left) + " " + preorder(t.Right)
//}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return traverse(s, t)
}

func traverse(s *TreeNode, t *TreeNode) bool {
	return s != nil && (equals(s, t) || traverse(s.Left, t) || traverse(s.Right, t))
}

func equals(x *TreeNode, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return x.Val == y.Val && equals(x.Left, y.Left) && equals(x.Right, y.Right)
}
