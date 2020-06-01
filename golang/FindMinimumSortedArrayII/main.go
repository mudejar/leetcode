package main

func main() {

}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		}
		return nums[0]
	}

	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left int, right int) int {
	if left >= right {
		return nums[left]
	}

	midpoint := left + (right-left)/2
	if nums[midpoint] < nums[right] {
		return search(nums, left, midpoint)
	}

	if nums[midpoint] > nums[right] {
		return search(nums, midpoint+1, right)
	}

	if nums[midpoint] == nums[right] {
		return search(nums, left, right-1)
	}

	return nums[0]
}
