package main

func main() {

}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	// This iteration will trickle down the existing matrix values and produce
	// the minimum cost available for every choice by recording the current
	// choices of values with whatever previous value can be used
	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}
	m := len(costs) - 1

	return min(min(costs[m][0], costs[m][1]), costs[m][2])
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
