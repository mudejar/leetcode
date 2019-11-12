package main

import "math"

func main() {

}

type MinStack struct {
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:(len(this.stack)-1)]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	for _, val := range this.stack {
		if val < min {
			min = val
		}
	}
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */