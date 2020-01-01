package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	size := getListSize(head)

	if size == 1 {
		head = nil
		return head
	}

	positionToRemove := size - n
	currentPosition := 0

	// Means we are removing the first element of the list.
	if positionToRemove == 0 {
		head = head.Next
		return head
	}

	tempNode := &ListNode{}
	current := head

	for currentPosition != positionToRemove {
		currentPosition++
		tempNode = current
		current = current.Next
	}

	if current.Next == nil {
		tempNode.Next = nil
	} else if current.Next != nil {
		tempNode.Next = current.Next
		current.Next = nil
		current = nil
	}

	return head
}

func getListSize(head *ListNode) int {
	current := head
	size := 0
	for current != nil {
		size++
		current = current.Next
	}
	return size
}
