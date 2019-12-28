package main

func main() {

}

func splitArray(nums []int, m int) int {
	// Find the highest single value in the array as well as sum total
	// Set left pointer to the single value, set right pointer to sum total
	left := 0; right := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}

	// start the binary search
	// starting point in the sum total of array, keep searching for a lower
	// value until there is nothing left to search
	answer := right
	for left <= right {
		midpoint := left + (right-left)/2
		sum := 0
		count := 1
		// loop across the whole nums array and add up a sum until it is higher
		// than the midpoint value you're examining. Once found, reset the sum
		// and increment the count value. The count value counts how many subarrays
		// are currently being accounted for.
		for i := 0; i < n; i++ {
			if sum + nums[i] > midpoint {
				count++
				sum = nums[i]
			} else {
				sum += nums[i]
			}
		}

		if count <= m {
			answer = min(answer, midpoint)
			right = midpoint - 1
		} else {
			left = midpoint + 1
		}
	}
	return answer
}

func min(x,y int) int {
	if x < y {
		return x
	}

	return y
}