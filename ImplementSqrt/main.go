package main

import(
	log "github.com/sirupsen/logrus"
	"math"
)

func main() {
	log.WithFields(log.Fields{
		"4": mySqrt(4),
		"8": mySqrt(8),
		"9": mySqrt(9),
		"2": mySqrt(2),
		"2147395610": mySqrt(2147395610),
	}).Println()
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