package main

func main() {

}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	diffs := map[int]int{}

	n := len(A)
	result := 0

	// Get all possible two-sums between C and D
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := C[i] + D[j]
			diffs[sum]++
		}
	}

	// Count the number of two-sums between A and B that equals to opposite of any two-sum between C and D
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := A[i] + B[j]
			result += diffs[-sum]
		}
	}

	return result
}

// var tuples int
//
// func fourSumCount(A []int, B []int, C []int, D []int) int {
// 	for _, val := range A {
// 		helper(val, B, C, D)
// 	}
// 	return tuples
// }
//
// func helper(prev int, arrs ...[]int) {
// 	if len(arrs) == 1 {
// 		for _, val := range arrs[0] {
// 			if prev+val == 0 {
// 				tuples++
// 			}
// 		}
// 	}
//
// 	arr := arrs[0]
// 	for _, val := range arr {
// 		helper(prev+val, arrs[1:]...)
// 	}
// }
