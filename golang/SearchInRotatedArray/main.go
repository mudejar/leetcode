package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("[]int{4,5,6,7,0,1,2}, target = 0", search([]int{4, 5, 6, 7, 0, 1, 2}, 0)).Println()
	log.WithField("[]int{1,3}, target = 0", search([]int{1, 3}, 0)).Println()
	log.WithField("[]int{5,1,3}, target = 5", search([]int{5, 1, 3}, 5)).Println()
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	rotationIndex := findRotationIndex(nums)
	var left, right int
	if rotationIndex == 0 {
		left = 0
		right = len(nums) - 1
	} else if nums[0] > target {
		left = rotationIndex
		right = len(nums) - 1
	} else {
		left = 0
		right = rotationIndex - 1
	}
	return binarySearch(nums, target, left, right)
}

func findRotationIndex(nums []int) int {
	rotationIndex := 0
	left := 0
	right := len(nums) - 1
	for left <= right {
		midpoint := (left + right) / 2
		if midpoint >= len(nums)-1 {
			return rotationIndex
		}
		if nums[midpoint] > nums[midpoint+1] {
			rotationIndex = midpoint + 1
			break
		} else {
			if nums[left] > nums[midpoint] {
				right = midpoint - 1
			} else {
				left = midpoint + 1
			}
		}
	}

	return rotationIndex
}

func binarySearch(arr []int, target int, left int, right int) int {
	if right-left <= 1 {
		if arr[left] == target {
			return left
		} else if arr[right] == target {
			return right
		} else {
			return -1
		}
	}

	midpoint := (left + right) / 2
	if arr[midpoint] < target {
		return binarySearch(arr, target, midpoint, right)
	} else if arr[midpoint] > target {
		return binarySearch(arr, target, left, midpoint)
	}

	return midpoint
}
