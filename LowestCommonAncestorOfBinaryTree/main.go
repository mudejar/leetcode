package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	arr []*TreeNode
}

func NewStack() *Stack {
	return &Stack{
		arr: []*TreeNode{},
	}
}

func (s *Stack) Push(node *TreeNode) {
	s.arr = append(s.arr, node)
}

func (s *Stack) Pop() *TreeNode {
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Empty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	nodeStack := NewStack()
	set := map[*TreeNode]struct{}{}

	search(root, p, nodeStack, nil)
	search(root, q, nil, &set)

	for !nodeStack.Empty() {
		elem := nodeStack.Pop()
		_, found := set[elem]
		if found {
			return elem
		}
	}

	return root
}

func search(root *TreeNode, target *TreeNode, nodeStack *Stack, set *map[*TreeNode]struct{}) *TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		populate(root, nodeStack, set)
		return root
	}

	if root.Val > target.Val {
		populate(root, nodeStack, set)
		return search(root.Left, target, nodeStack, set)
	}

	populate(root, nodeStack, set)
	return search(root.Right, target, nodeStack, set)
}

func populate(node *TreeNode, nodeStack *Stack, set *map[*TreeNode]struct{}) {
	if nodeStack != nil {
		nodeStack.Push(node)
	}

	if set != nil {
		(*set)[node] = struct{}{}
	}
}
