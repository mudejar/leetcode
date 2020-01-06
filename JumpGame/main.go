package main

func main() {

}

// Dynamic programming bottom up solution: Time: O(n^2) Space: O(n)
//
// const (
// 	GOOD = iota
// 	BAD
// 	UNKNOWN
// )
//
// func canJump(nums []int) bool {
// 	memo := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		memo[i] = UNKNOWN
// 	}
// 	memo[len(memo)-1] = GOOD
//
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		furthestJump := min(nums[i]+i, len(nums)-1)
// 		for j := i + 1; j <= furthestJump; j++ {
// 			if memo[j] == GOOD {
// 				memo[i] = GOOD
// 				break
// 			}
// 		}
// 	}
//
// 	return memo[0] == GOOD
// }
//
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
//
// 	return b
// }

// Greedy solution
func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
