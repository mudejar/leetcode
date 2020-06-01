package main

import (
	"math"
	"sort"
)

func main() {

}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	left := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < left {
			left = nums[i] - nums[i-1]
		}
	}
	right := nums[len(nums)-1] - nums[0]

	for left < right {
		mid := (left + right) / 2
		count := 0

		for i, j := 0, 1; i < len(nums); i++ {
			for j < len(nums) && nums[j]-nums[i] <= mid {
				j++
			}
			count += j - i - 1
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
