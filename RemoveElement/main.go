package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
	"3 2 2 3": removeElement([]int{3,2,2,3}, 3),
	"0 1 2 2 3 0 4 2": removeElement([]int{0,1,2,2,3,0,4,2}, 2),
	"1": removeElement([]int{1}, 1),
	}).Println()
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}