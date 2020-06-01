package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for i := head; i != nil; i = i.Next {
		arr = append(arr, i.Val)
	}

	length := len(arr)
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-i-1] {
			return false
		}
	}

	return true
}
