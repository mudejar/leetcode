package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"() true": isValid("()"),
		"()[]{} true": isValid("()[]{}"),
		"(] false": isValid("(]"),
		"([)] false": isValid("([)]"),
		"{[]} true": isValid("{[]}"),
		"[": isValid("["),
	}).Println()
}

func isValid(s string) bool {
	openBrackets := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	closeBrackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var brackets = ""
	for _, letter := range s {
		_, exists := openBrackets[string(letter)]
		if exists {
			// push to the string of brackets
			brackets += string(letter)
		}

		counterpart, exists := closeBrackets[string(letter)]
		if exists {
			if len(brackets) < 1 {
				return false
			}
			// pop the last element in the string
			pop := pop(&brackets)
			if pop != counterpart {
				return false
			}
		}
	}

	if len(brackets) > 0 {
		return false
	}

	return true
}

func pop(s *string) string{
	element := string((*s)[len(*s)-1])
	(*s) = (*s)[:len(*s)-1]
	return element
}