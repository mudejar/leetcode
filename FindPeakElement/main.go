package main

func main() {

}

func findPeakElement(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums)-1
	}

	left, right := 1, len(nums)-2
	return findPeak(nums, left, right)
}

func findPeak(nums []int, left int, right int) int {
	midpoint := (right + left) / 2
	if nums[midpoint] > nums[midpoint+1] && nums[midpoint] > nums[midpoint-1] {
		return midpoint
	}

	if nums[midpoint] < nums[midpoint+1] {
		return findPeak(nums, midpoint+1, right)
	}

	if nums[midpoint] < nums[midpoint-1] {
		return findPeak(nums, left, midpoint-1)
	}
	return 0
}
