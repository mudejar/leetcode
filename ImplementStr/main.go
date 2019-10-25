package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		`haystack = "hello", needle = "ll"`: strStr("hello", "ll"),
		`"aaaaa", needle = "bba"`: strStr("aaaaa", "bba"),
		`"AABAACAADAABAABA" needle="AABA"`: strStr("AABAACAADAABAABA", "AABA"),
		`haystack="ABXABABXAB" needle="ABXAB"`: strStr("ABXABABXAB", "ABXAB"),
	}).Println()
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	return KMPSearch(needle, haystack)
}

// Knuth-Morris-Pratt (KMP) in Golang
func KMPSearch(needle, haystack string) int {
	needleLength := len(needle)
	haystackLength := len(haystack)

	// create lps[] that will hold the longest prefix suffix
	// values for pattern
	lpsArr := getLPS(needle)

	j:=0; i:=0
	for i < haystackLength {
		if (needle[j] == haystack[i]) {
			j++
			i++
		}

		if j == needleLength {
			return i-j
		} else if i < haystackLength && needle[j] != haystack[i] {
			// Do not match lps[0..lps[j-1]] characters,
			// they will match anyway
			if j != 0 {
				j = lpsArr[j-1]
			} else {
				i = i + 1
			}
		}
	}

	return -1
}

func getLPS(needle string) []int {
	k := 0
	m := len(needle)
	lps := make([]int, m)
	lps[0] = 0
	// the loop calculates lps[i] for i = 1 to M-1
	for i := 1; i < m; {
		if needle[i] == needle[k] {
			k++
			lps[i] = k
			i++
		} else { // needle[i] != needle[k]
			// This is tricky. Consider the example.
			// AAACAAAA and i = 7. The idea is similar
			// to search step.
			if k != 0 {
				k = lps[k-1]

				// Also, note that we do not increment
				// i here
			} else { // if (k == 0)
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}