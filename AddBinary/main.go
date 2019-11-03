package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	log.WithFields(log.Fields{
		`a = "11", b = "1"`: addBinary("11", "1"),
		`a = "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", b = "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"`: addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"),
	}).Println()
}

//func addBinary(a string, b string) string {
//	if len(a) < len(b) {
//		a, b = b, a
//	}
//
//	carry := 0
//	s := ""
//	zeroes := ""
//
//	for i := 0; i < len(a)-len(b); i++ {
//		zeroes += "0"
//	}
//
//	b = zeroes + b
//
//	for i := len(a) - 1; i >= 0; i-- {
//		sum := carry
//		if string(a[i]) == "1" {
//			sum++
//		}
//
//		if string(b[i]) == "1" {
//			sum++
//		}
//
//		carry = sum / 2
//		sum = sum % 2
//		s = strconv.Itoa(sum) + s
//	}
//
//	if carry > 0 {
//		s = "1" + s
//	}
//
//	return s
//}

// Bit operations solution
func addBinary(a string, b string) string {
	x, _ := strconv.ParseInt(a, 2, 64)
	y, _ := strconv.ParseInt(b, 2, 64)

	for y != 0 {
		answer := x^y
		carry := (x & y) << 1
		x = answer
		y = carry
	}

	return strconv.FormatInt(x, 2)
}