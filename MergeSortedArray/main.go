package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	list := []int{1,2,3,0,0,0}
	merge(list, 3, []int{2,5,6},3)
	log.WithFields(log.Fields{
		`nums1 = [1,2,3,0,0,0], m = 3
		 nums2 = [2,5,6],       n = 3`: list,
	})
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}