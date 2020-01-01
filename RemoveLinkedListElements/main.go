package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode = nil
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		if cur.Val == val {
			if prev == nil { // the beginning of the list
				head = head.Next
			} else {
				prev.Next = cur.Next
			}
			cur = cur.Next
		} else {
			temp := cur
			cur = cur.Next
			prev = temp
		}
	}

	return head
}
