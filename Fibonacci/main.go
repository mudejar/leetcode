package main

func main() {

}

func fib(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	dp := []int{0,1}
	for i := 2; i <= N; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}

	return dp[N]
}