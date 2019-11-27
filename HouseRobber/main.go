package main

func main() {
	rob([]int{2, 1, 1, 2})
}

func rob(nums []int) int {
	prevMax := 0
	currMax := 0

	for _, num := range nums {
		temp := currMax
		currMax = max(prevMax+num, currMax)
		prevMax = temp
	}

	return currMax
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
