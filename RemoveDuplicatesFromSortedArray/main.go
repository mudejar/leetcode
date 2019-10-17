package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"[1,1,2]": removeDuplicates([]int{1,1,2}),
		"[0,0,1,1,1,2,2,3,3,4]": removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}),
	}).Println()
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j:=0; j<len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}