package main

import (
	"strings"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		`race a car`: isPalindrome("race a car"),
		`0P`: isPalindrome("0P"),
	}).Println()
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	filteredStr := ""
	for _, val := range s {
		if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') || (val >= '0' && val <= '9') {
			filteredStr += strings.ToLower(string(val))
		}
	}

	i := 0
	j := len(filteredStr)-1

	for i < len(filteredStr) && j > 0 {
		if filteredStr[i] != filteredStr[j] {
			return false
		}
		i++; j--
	}

	return true
}