package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("[]int{5,7,7,8,8,10}, 8", searchRange([]int{5, 7, 7, 8, 8, 10}, 8)).Println()
	log.WithField("[]int{1,2,3}, 1", searchRange([]int{1, 2, 3}, 1)).Println()
}

func searchRange(nums []int, target int) []int {
	arr := []int{-1, -1}
	if len(nums) == 0 {
		return arr
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return arr
	}

	if len(nums) == 2 {
		return []int{getLeftBound(nums, target, 0, len(nums)-1), getRightBound(nums, target, 0, len(nums)-1)}
	}

	pivot := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		midpoint := (right + left) / 2
		if nums[midpoint] == target {
			pivot = midpoint
			break
		} else if nums[midpoint] < target {
			left = midpoint + 1
		} else if nums[midpoint] > target {
			right = midpoint - 1
		}
	}

	if pivot == -1 {
		return []int{-1, -1}
	}

	return []int{getLeftBound(nums, target, 0, pivot), getRightBound(nums, target, pivot, len(nums)-1)}
}

func getLeftBound(nums []int, target int, left int, right int) int {
	if left == right-1 || left == right {
		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		return -1
	}

	midpoint := (right + left) / 2
	if nums[midpoint] == target && nums[midpoint-1] != target {
		return midpoint
	}

	if nums[midpoint] != target {
		return getLeftBound(nums, target, midpoint+1, right)
	}

	if nums[midpoint] == target {
		return getLeftBound(nums, target, left, midpoint-1)
	}

	return -1
}

func getRightBound(nums []int, target int, left int, right int) int {
	if left+1 == right || left == right {
		if nums[right] == target {
			return right
		}

		if nums[left] == target {
			return left
		}

		return -1
	}

	midpoint := (right + left) / 2
	if nums[midpoint] == target && nums[midpoint+1] != target {
		return midpoint
	}

	if nums[midpoint] != target {
		return getRightBound(nums, target, left, midpoint-1)
	}

	if nums[midpoint] == target {
		return getRightBound(nums, target, midpoint+1, right)
	}

	return -1
}
