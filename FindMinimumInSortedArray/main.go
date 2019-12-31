package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("[4,5,6,7,0,1,2]", findMin([]int{4, 5, 6, 7, 0, 1, 2})).Println()
	log.WithField("[1,2]", findMin([]int{1, 2})).Println()
	log.WithField("[1,2,3]", findMin([]int{1, 2, 3})).Println()
	log.WithField("[3,1,2]", findMin([]int{3, 1, 2})).Println()
	log.WithField("[5,1,2,3,4]", findMin([]int{5, 1, 2, 3, 4})).Println()
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

	return binarySearch(nums, 0, len(nums)-1)
}

//func binarySearch(nums []int, left int, right int) int {
//	if right == left {
//		if right == len(nums)-1 && nums[right] > nums[0] {
//			return nums[0]
//		}
//		return nums[left]
//	}
//
//	midpoint := (left + right) / 2
//	if nums[midpoint] > nums[midpoint+1] {
//		return nums[midpoint+1]
//	}
//
//	if nums[midpoint] < nums[midpoint-1] {
//		return nums[midpoint]
//	}
//
//	if nums[midpoint] < nums[0] {
//		return binarySearch(nums, left, midpoint-1)
//	}
//
//	return binarySearch(nums, midpoint+1, right)
//}

func binarySearch(nums []int, left int, right int) int {
	if right == left {
		if right == len(nums)-1 && nums[right] > nums[0] {
			return nums[0]
		}
		return nums[left]
	}

	midpoint := (left + right) / 2
	if nums[midpoint] > nums[midpoint+1] {
		return nums[midpoint+1]
	}

	if nums[midpoint] < nums[midpoint-1] {
		return nums[midpoint]
	}

	if nums[midpoint] < nums[0] {
		return binarySearch(nums, left, midpoint-1)
	}

	return binarySearch(nums, midpoint+1, right)
}
