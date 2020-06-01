package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// the iterative option
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for head.Next != nil {
		p := head.Next
		head.Next = p.Next
		p.Next = cur
		cur = p
	}

	return cur
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
//
// 	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// 	node := reverseList(head.Next) // node = 5
// 	head.Next.Next = head          // 1 -> 2 -> 3 -> 4 -> 5 -> 4
// 	head.Next = nil                // 1 -> 2 -> 3 -> 4 x 5 -> 4
// 	return node                    // 5 -> 4
// }

// 1 -> 2 -> 3 -> 4 x 5 -> 4
// 1 -> 2 -> 3 -> 4 x 5 -> 4 -> 3
// 1 -> 2 -> 3 x 4 x 5 -> 4 -> 3
