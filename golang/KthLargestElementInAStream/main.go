package main

func main() {

}

type TreeNode struct {
	Val   int
	Cnt   int
	Left  *TreeNode
	Right *TreeNode
}

type KthLargest struct {
	Root *TreeNode
	MK   int
}

func (k *KthLargest) searchKth(root *TreeNode, K int) int {
	// m = the size of the right subtree
	var m int
	if root.Right != nil {
		m = root.Right.Cnt
	} else {
		m = 0
	}

	if K == m+1 {
		return root.Val
	}

	if K <= m {
		// find the kth largest in the right subtree
		return k.searchKth(root.Right, K)
	} else {
		return k.searchKth(root.Left, K-m-1)
	}
}

func Constructor(k int, nums []int) KthLargest {
	var root *TreeNode
	for i := 0; i < len(nums); i++ {
		root = insertNode(root, nums[i])
	}

	return KthLargest{
		Root: root,
		MK:   k,
	}
}

func (this *KthLargest) Add(val int) int {
	this.Root = insertNode(this.Root, val)
	return this.searchKth(this.Root, this.MK)
}

func insertNode(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: num,
			Cnt: 1,
		}
	}

	if root.Val < num {
		root.Right = insertNode(root.Right, num)
	} else {
		root.Left = insertNode(root.Left, num)
	}
	root.Cnt++
	return root
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
