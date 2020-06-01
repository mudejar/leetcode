package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("Answer", intersect([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4})).Println()
}

func intersect(nums1 []int, nums2 []int) []int {
	nums1, nums2 = orderByLength(nums1, nums2)

	m := map[int]int{}
	for _, val := range nums1 {
		m[val]++
	}

	log.WithField("m", m).Println() // debug
	log.WithFields(log.Fields{
		"nums1": nums1,
		"nums2": nums2,
	}).Println() // debug

	end := 0
	for k := 0; k < len(nums2); k++ {
		val := m[nums2[k]]
		if val > 0 {
			nums1[end] = nums2[k]
			m[nums2[k]]--
			end++
		}
	}

	return nums1[:end]
}

func orderByLength(nums1, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums1, nums2
	}

	return nums2, nums1
}
