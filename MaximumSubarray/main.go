package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"[-2,1,-3,4,-1,2,1,-5,4]": maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}),
	}).Println()
}

// Greedy solution
func maxSubArray(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}