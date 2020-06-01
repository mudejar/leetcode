package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	nodes := map[*ListNode]struct{}{}
	for ; head != nil; head = head.Next {
		_, exists := nodes[head]
		if exists {
			return head
		}
		nodes[head] = struct{}{}
	}

	return head
}
