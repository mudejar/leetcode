package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"20": isPerfectSquare(20),
	}).Println()
}

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}

	return search(num, 0, num)
}

func search(target, left, right int) bool {
	if left >= right {
		if left*left == target {
			return true
		}

		return false
	}

	midpoint := (left+right)/2
	square := midpoint*midpoint
	if square == target {
		return true
	}

	if square < target {
		return search(target, midpoint+1, right)
	}

	if square > target {
		return search(target, left, midpoint-1)
	}

	return false
}