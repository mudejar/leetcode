package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Hello World":  lengthOfLastWord("Hello World"),
		"Hello World ": lengthOfLastWord("Hello World "),
		"a":            lengthOfLastWord("a"),
	}).Println()
}

//func lengthOfLastWord(s string) int {
//	if len(s)== 0 {
//		return 0
//	}
//
//	wordLength := 0
//	newWord := true
//	for _, letter := range s {
//		if letter == ' ' {
//			newWord = true
//		} else {
//			if newWord == true {
//				wordLength = 1
//				newWord = false
//			} else {
//				wordLength++
//			}
//		}
//	}
//
//	return wordLength
//}

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	l := len(s)
	for l > 0 && s[l-1] == ' ' {
		l--
		s = s[:l]
	}

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		length++
	}
	return length
}
