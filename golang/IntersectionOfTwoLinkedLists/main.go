package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	for curA != nil && curB != nil {
		curA = curA.Next
		curB = curB.Next
	}

	for curA != nil {
		headA = headA.Next
		curA = curA.Next
	}

	for curB != nil {
		headB = headB.Next
		curB = curB.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
