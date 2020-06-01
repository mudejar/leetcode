package main

func main() {

}

func maxArea(height []int) int {
	totalWater := 0
	l, r := 0, len(height)-1

	for l < r {
		width := r - l
		ht := min(height[l], height[r])
		totalWater = max(totalWater, width*ht)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return totalWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
