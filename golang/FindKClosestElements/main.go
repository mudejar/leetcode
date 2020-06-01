package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"[]int{1,2,3,4,5}, k = 4, x = 3": findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3),
		//"[]int{0,1,1,1,2,3,6,7,8,9}, k = 9, x = 4": findClosestElements([]int{0,1,1,1,2,3,6,7,8,9}, 9, 4),
		"[]int{1,1,2,3,3,3,4,6,8,8}, k = 6, x = 1": findClosestElements([]int{1, 1, 2, 3, 3, 3, 4, 6, 8, 8}, 6, 1),
	}).Println()
}

func findClosestElements(arr []int, k int, x int) []int {
	alen := len(arr)
	if alen == 0 {
		return []int{}
	}

	if alen == 1 {
		return arr
	}

	if x < arr[0] {
		return arr[:k]
	}

	if x > arr[alen-1] {
		return arr[(alen - 1 - k):]
	}

	index := searchForIndex(arr, x, 0, alen-1)

	l := index
	r := index + 1
	answer := []int{}
	for len(answer) < k {
		if r > alen-1 {
			answer = append([]int{arr[l]}, answer...)
			l--
		} else if l < 0 {
			answer = append(answer, arr[r])
			r++
		} else {
			if absoluteValue(x, arr[l]) <= absoluteValue(x, arr[r]) {
				answer = append([]int{arr[l]}, answer...)
				l--
			} else {
				answer = append(answer, arr[r])
				r++
			}
		}
	}
	return answer
}

func searchForIndex(arr []int, target int, left int, right int) int {
	if left == right-1 || left == right {
		if arr[left] == target {
			return left
		}

		return right
	}

	midpoint := (left + right) / 2
	if arr[midpoint] == target && arr[midpoint-1] != target {
		return midpoint
	}

	if arr[midpoint] < target {
		return searchForIndex(arr, target, midpoint+1, right)
	}

	return searchForIndex(arr, target, left, midpoint-1)
}

func absoluteValue(target int, value int) int {
	absVal := target - value
	if absVal < 0 {
		absVal = -1 * absVal
	}
	return absVal
}
