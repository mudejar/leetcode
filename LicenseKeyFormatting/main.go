package main

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.WithFields(log.Fields{
		`"2-5g-3-J" 2`: licenseKeyFormatting("2-5g-3-J", 2),
	}).Println()
}

func licenseKeyFormatting(S string, K int) string {
	stringWithNoDashes := ""
	for i := len(S) - 1; i >= 0; i-- {
		if string(S[i]) != "-" {
			stringWithNoDashes = string(S[i]) + stringWithNoDashes
		}
	}

	formattedString := ""
	k := 0
	for i := len(stringWithNoDashes) - 1; i >= 0; i-- {
		formattedString = string(stringWithNoDashes[i]) + formattedString
		k++
		if k == K && i != 0 {
			formattedString = "-" + formattedString
			k = 0
		}
	}

	return strings.ToUpper(formattedString)
}
