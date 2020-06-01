package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return true
		}

		if p2.Next == nil {
			return false
		}

		p2 = p2.Next.Next
		p1 = p1.Next
	}

	return false
}
