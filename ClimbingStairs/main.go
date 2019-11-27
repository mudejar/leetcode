package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"2":  climbStairs(2),
		"3":  climbStairs(3),
		"45": climbStairs(45),
	}).Println()
}

// Dynamic programming solution
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := []int{0, 1, 2}
	for i := 3; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[n]
}
