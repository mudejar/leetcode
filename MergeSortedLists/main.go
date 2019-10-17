package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	log.WithFields(log.Fields{
		"1->2->4, 1->3->4": mergeTwoLists(l1, l2),
	}).Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 == nil {
//		return nil
//	}
//	if l2 == nil {
//		return &ListNode{
//			Val:  l1.Val,
//			Next: mergeTwoLists(l1.Next, nil),
//		}
//	}
//	if l1 == nil {
//		return &ListNode{
//			Val:  l2.Val,
//			Next: mergeTwoLists(nil, l2.Next),
//		}
//	}
//
//	if l1.Val < l2.Val {
//		return &ListNode{
//			Val:  l1.Val,
//			Next: mergeTwoLists(l1.Next, l2),
//		}
//	} else {
//		return &ListNode{
//			Val:  l2.Val,
//			Next: mergeTwoLists(l1, l2.Next),
//		}
//	}
//}

// This is the faster solution
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}