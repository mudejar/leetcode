package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	node := reverseList(head.Next) // node = 5
	head.Next.Next = head          // 1 -> 2 -> 3 -> 4 -> 5 -> 4
	head.Next = nil                // 1 -> 2 -> 3 -> 4 -> nil -> 4
	return node                    // 5 -> 4
}
