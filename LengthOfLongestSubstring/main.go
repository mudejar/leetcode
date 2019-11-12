package main

import "math"

func main() {

}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	charToIndeces := map[uint8]int{}
	longestLength := 0; i := 0
	for j := 0; j < length; j++ {
		_, exists := charToIndeces[s[j]]
		if exists {
			i = int(math.Max(float64(charToIndeces[s[j]]), float64(i)))
		}
		longestLength = int(math.Max(float64(longestLength), float64(j-i+1)))
		charToIndeces[s[j]] = j+1
	}
	return longestLength
}
