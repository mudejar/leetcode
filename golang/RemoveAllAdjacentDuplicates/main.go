package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type stack struct {
	arr []int32
}

func (s *stack) Len() int {
	return len(s.arr)
}

func (s *stack) Pop() *int32 {
	val := s.arr[len(s.arr)]
	s.arr = s.arr[:len(s.arr)-1]
	return &val
}

func (s *stack) Push(val int32) {
	s.arr = append(s.arr, val)
}

func (s *stack) Peek() *int32 {
	if len(s.arr) == 0 {
		return nil
	}

	return &s.arr[len(s.arr)-1]
}

func main() {
	log.WithField(
		"Answer", removeDuplicates("abbaca"),
	).Println()
}

//func removeDuplicates(S string) string {
//	length := len(S)
//	j := 1
//	for j < length {
//		if S[j] == S[j-1] {
//			S = S[:j-1] + S[j+1:]
//			length = len(S)
//			j = 1
//		} else {
//			j++
//		}
//	}
//
//	return S
//}

func removeDuplicates(S string) string {
	st := stack{make([]int32, 0)}

	for _, val := range S {
		if st.Peek() != nil {
			if *st.Peek() == val {
				st.Pop()
			} else {
				st.Push(val)
			}
		} else {
			st.Push(val)
		}
	}

	str := ""
	for st.Len() > 0 {
		str = str + fmt.Sprintf("%d", st.Pop())
	}

	return str
}
