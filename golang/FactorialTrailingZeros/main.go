package main

func main() {

}

func trailingZeroes(n int) int {
	pairs := 0
	denominator := 5
	for denominator <= n {
		pairs += n / denominator
		denominator *= 5
	}

	return pairs
}

// func trailingZeroes(n int) int {
// 	pairs := 0
// 	for i := 5; i <= n; i += 5 {
// 		remainingI := i
// 		for remainingI%5 == 0 {
// 			pairs++
// 			remainingI /= 5
// 		}
// 	}
//
// 	return pairs
// }
