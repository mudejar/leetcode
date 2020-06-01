package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		`["test.email+alex@leetcode.com", "test.email@leetcode.com"]`: numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}),
	}).Println()
}

func numUniqueEmails(emails []string) int {
	existingEmails := map[string]struct{}{}
	total := 0
	for _, email := range emails {
		filteredEmail := filter(email)
		_, exists := existingEmails[filteredEmail]
		if !exists {
			existingEmails[filteredEmail] = struct{}{}
			total++
		}
	}

	return total
}

func filter(email string) string {
	filteredEmail := ""
	for i, letter := range email {
		switch string(letter) {
		case ".":
			continue
		case "+":
			j := i
			for ; j < len(email); j++ {
				if string(email[j]) == "@" {
					break
				}
			}
			return filteredEmail + email[j:]
		case "@":
			return filteredEmail + email[i:]
		default:
			filteredEmail += string(letter)
		}
	}
	return filteredEmail
}
