package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(5)
	stack.Push(1)
	stack.Push(5)
	fmt.Println("stack.Top(): ", stack.Top())
	fmt.Println("stack.PopMax(): ", stack.PopMax())
	fmt.Println("stack.Top(): ", stack.Top())
	fmt.Println("stack.PeekMax(): ", stack.PeekMax())
	fmt.Println("stack.Pop(): ", stack.Pop())
	fmt.Println("stack.Top(): ", stack.Top())
}

type Node struct {
	Val      int
	MaxValue int
	Next     *Node
}

type MaxStack struct {
	Head *Node
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	newHead := &Node{
		Val:  x,
		Next: this.Head,
	}

	if this.Head != nil {
		if x > this.Head.MaxValue {
			newHead.MaxValue = x
		} else {
			newHead.MaxValue = this.Head.MaxValue
		}
	} else {
		newHead.MaxValue = x
	}

	this.Head = newHead
}

func (this *MaxStack) Pop() int {
	val := this.Head.Val
	this.Head = this.Head.Next
	return val
}

func (this *MaxStack) Top() int {
	return this.Head.Val
}

func (this *MaxStack) PeekMax() int {
	return this.Head.MaxValue
}

func (this *MaxStack) PopMax() int {
	prev := this.Head

	if prev.Val == prev.MaxValue {
		val := prev.MaxValue
		this.Head = this.Head.Next
		return val
	}

	for node := prev.Next; node != nil; node = node.Next {
		if node.Val == node.MaxValue {
			prev.Next = node.Next
			return node.Val
		}
	}

	return 0
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
