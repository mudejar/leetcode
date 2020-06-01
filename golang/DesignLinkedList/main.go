package main

func main() {

}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.Head
	for i := 0; i <= index && node != nil; i++ {
		if i == index {
			return node.Val
		}

		node = node.Next
	}

	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{
		Val:  val,
		Next: this.Head,
	}
	this.Head = newHead
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var prev *Node = nil
	cur := this.Head
	for cur != nil {
		tmp := cur
		cur = cur.Next
		prev = tmp
	}

	prev.Next = &Node{
		Val: val,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Head == nil && index == 0 {
		this.Head = &Node{
			Val: val,
		}
		return
	}

	if index == 0 {
		newHead := &Node{
			Val:  val,
			Next: this.Head,
		}
		this.Head = newHead
		return
	}

	if index < 0 {
		return
	}

	var prev *Node = nil
	cur := this.Head
	i := 0
	for ; cur != nil; i++ {
		if i == index {
			prev.Next = &Node{
				Val:  val,
				Next: cur,
			}
			return
		}
		tmp := cur
		cur = cur.Next
		prev = tmp
	}

	if i < index {
		return
	}

	prev.Next = &Node{
		Val: val,
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Head == nil {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	var prev *Node = nil
	cur := this.Head
	for i := 0; cur != nil; i++ {
		if i == index {
			prev.Next = cur.Next
			return
		}
		tmp := cur
		cur = cur.Next
		prev = tmp
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
