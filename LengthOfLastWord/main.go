package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"Hello World": lengthOfLastWord("Hello World"),
		"Hello World ": lengthOfLastWord("Hello World "),
		"a": lengthOfLastWord("a"),
	}).Println()
}

func lengthOfLastWord(s string) int {
	if len(s)== 0 {
		return 0
	}

	wordLength := 0
	newWord := true
	for _, letter := range s {
		if letter == ' ' {
			newWord = true
		} else {
			if newWord == true {
				wordLength = 1
				newWord = false
			} else {
				wordLength++
			}
		}
	}

	return wordLength
}