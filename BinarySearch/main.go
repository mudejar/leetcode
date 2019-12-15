package BinarySearch

func main() {

}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return searchHelper(nums, target, 0, len(nums)-1)
}

func searchHelper(nums []int, target int, left int, right int) int {
	if right-left <= 1 {
		if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		} else {
			return -1
		}
	}

	midpoint := (right-left)/2 + left

	if nums[midpoint] > target {
		return searchHelper(nums, target, left, midpoint)
	} else if nums[midpoint] < target {
		return searchHelper(nums, target, midpoint, right)
	}

	return midpoint
}
