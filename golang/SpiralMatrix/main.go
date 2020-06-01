package main

import "fmt"

// # Given a matrix of m x n elements (m rows, n columns),
// # return all elements of the matrix in spiral order.

// # Input
// # [
// #  [ 1, 2, 3 ],
// #  [ 4, 5, 6 ],
// #  [ 7, 8, 9 ]
// # ]
// # Output: [1,2,3,6,9,8,7,4,5]

// # Input:
// # [
// #   [1, 2, 3, 4],
// #   [5, 6, 7, 8],
// #   [9,10,11,12]
// # ]
// # Output: [1,2,3,4,8,12,11,10,9,5,6,7]

type Node struct {
	Val       int
	remaining int
	matrix    [][]int
	r         int
	c         int
}

func newNode(r int, c int, remaining int, m [][]int) *Node {
	return &Node{
		Val:       m[r][c],
		remaining: remaining,
		matrix:    m,
		r:         r,
		c:         c,
	}
}

func (n *Node) isEnd() bool {
	if n.remaining == 0 {
		return true
	}
	return false
}

func (n *Node) serialize() string {
	return string(n.r) + string(n.c)
}

func (f *Fsm) GoEast() {
	f.Location = newNode(f.Location.r, f.Location.c+1, f.Location.remaining-1, f.Location.matrix)
}

func (f *Fsm) GoWest() {
	f.Location = newNode(f.Location.r, f.Location.c-1, f.Location.remaining-1, f.Location.matrix)
}

func (f *Fsm) GoSouth() {
	f.Location = newNode(f.Location.r-1, f.Location.c, f.Location.remaining-1, f.Location.matrix)
}

func (f *Fsm) GoNorth() {
	f.Location = newNode(f.Location.r+1, f.Location.c, f.Location.remaining-1, f.Location.matrix)
}

func helper(r int, c int, remaining int, matrix [][]int, visited map[string]struct{}) bool {
	next := newNode(r, c, remaining, matrix)
	_, ok := visited[next.serialize()]
	if ok {
		return false
	}
	return true
}

type Fsm struct {
	Visited  map[string]struct{}
	Location *Node
	Matrix   [][]int
}

func newFsm(matrix [][]int) *Fsm {
	if len(matrix) == 0 {
		return nil
	}

	firstNode := newNode(0, 0, len(matrix)*len(matrix[0]), matrix)
	return &Fsm{
		Visited: map[string]struct{}{
			firstNode.serialize(): struct{}{},
		},
		Location: firstNode,
	}
}

func (f *Fsm) CanGoNorth() bool {
	if f.Location.r+1 >= len(f.Location.matrix[0]) {
		return false
	}

	return helper(f.Location.r+1,
		f.Location.c,
		f.Location.remaining-1,
		f.Location.matrix,
		f.Visited,
	)
}

func (f *Fsm) CanGoEast() bool {
	if f.Location.c+1 >= len(f.Location.matrix) {
		return false
	}

	return helper(f.Location.r,
		f.Location.c+1,
		f.Location.remaining-1,
		f.Location.matrix,
		f.Visited,
	)
}

func (f *Fsm) CanGoWest() bool {
	if f.Location.c-1 < 0 {
		return false
	}

	return helper(f.Location.r,
		f.Location.c-1,
		f.Location.remaining-1,
		f.Location.matrix,
		f.Visited,
	)
}

func (f *Fsm) CanGoSouth(visited map[string]struct{}) bool {
	if f.Location.r-1 < 0 {
		return false
	}

	return helper(f.Location.r-1,
		f.Location.c,
		f.Location.remaining-1,
		f.Location.matrix,
		f.Visited,
	)
}

func (f *Fsm) generateSpiralArray() []int {
	arr := []int{f.Location.Val}
	for true {
		for true {
			if f.CanGoEast() {
				f.GoEast()
				arr = append(arr, f.Location.Val)
				fmt.Println(arr) // debug
				f.Visited[f.Location.serialize()] = struct{}{}
				if f.Location.isEnd() {
					return arr
				}
			} else {
				break
			}
		}

		for true {
			if f.CanGoSouth(f.Visited) {
				f.GoSouth()
				arr = append(arr, f.Location.Val)
				f.Visited[f.Location.serialize()] = struct{}{}
				if f.Location.isEnd() {
					return arr
				}
			} else {
				break
			}
		}

		for true {
			if f.CanGoWest() {
				f.GoWest()
				arr = append(arr, f.Location.Val)
				f.Visited[f.Location.serialize()] = struct{}{}
				if f.Location.isEnd() {
					return arr
				}
			} else {
				break
			}
		}

		for true {
			if f.CanGoNorth() {
				f.GoNorth()
				arr = append(arr, f.Location.Val)
				f.Visited[f.Location.serialize()] = struct{}{}
				if f.Location.isEnd() {
					return arr
				}
			} else {
				break
			}
		}
	}

	return arr
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	fsm := newFsm(matrix)
	return fsm.generateSpiralArray()
}

func main() {

}

// func spiralOrder(matrix [][]int) []int {
// 	if len(matrix) == 0 {
// 		return []int{}
// 	}
//
// 	N := 0
// 	E := len(matrix[0]) - 1
// 	S := len(matrix) - 1
// 	W := -1
//
// 	arr := []int{}
// 	total := len(matrix[0]) * len(matrix)
// 	spiraling := true
// 	for spiraling {
// 		spiraling = false
// 		for c := W + 1; c <= E; c++ {
// 			spiraling = true
// 			arr = append(arr, matrix[N][c])
// 			total--
// 		}
// 		if total == 0 {
// 			break
// 		}
// 		W++
//
// 		for r := N + 1; r <= S; r++ {
// 			spiraling = true
// 			arr = append(arr, matrix[r][E])
// 			total--
// 		}
// 		if total == 0 {
// 			break
// 		}
// 		N++
//
// 		for c := E - 1; c >= W; c-- {
// 			spiraling = true
// 			arr = append(arr, matrix[S][c])
// 			total--
// 		}
// 		if total == 0 {
// 			break
// 		}
// 		E--
//
// 		for r := S - 1; r >= N; r-- {
// 			spiraling = true
// 			arr = append(arr, matrix[r][W])
// 			total--
// 		}
// 		if total == 0 {
// 			break
// 		}
// 		S--
// 	}
//
// 	return arr
// }
