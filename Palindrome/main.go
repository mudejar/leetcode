package main

import (
	log "github.com/sirupsen/logrus"
	"math"
)

func main() {
	log.WithField("12331 should be false", isPalindrome(12331)).Println()
	log.WithField("121 should be true", isPalindrome(121)).Println()
	log.WithField("1221 should be true", isPalindrome(1221)).Println()
}

func findLength(x int) int {
	return int(math.Log10(float64(x))+1)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length := findLength(x)
	isEven := (length % 2) == 0
	arr := []int{}
	for i := 0; i < length/2; i++ {
		pop := x % 10
		x /= 10
		arr = append(arr, pop)
	}

	if !isEven {
		x /= 10
	}

	for i := 0; i < length/2; i++ {
		pop := x % 10
		x /= 10
		if pop != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}