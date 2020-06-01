package main

import (
	log "github.com/sirupsen/logrus"
	"math"
)

func main() {
	log.WithField("mySqrt", mySqrt(2)).Println()
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := int(math.Exp(0.5 * math.Log(float64(x))))
	right := left + 1

	if (right * right) > x {
		return left
	}

	return right
}
