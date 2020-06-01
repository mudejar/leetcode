package main

import (
	"sort"
	"strings"
)

//Reorder the logs so that all of the letter-logs come before any digit-log.
// The letter-logs are ordered lexicographically ignoring identifier, with
// the identifier used in case of ties.  The digit-logs should be put in their
// original order.
func main() {

}

func reorderLogFiles(logs []string) []string {
	letterLogs := make([]string, 0)
	digitLogs := make([]string, 0)

	// reorder the log files so that one slice is the digit logs and the other
	// holds the letter logs
	for _, log := range logs {
		if isDigit(log) {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	sort.Slice(letterLogs, func(i int, j int) bool { // This will perform a switch operation with a custom less function passed in
		ii := strings.Index(letterLogs[i], " ") // the index of the first instance of a space character
		ji := strings.Index(letterLogs[j], " ") // the index of the first instance of a space character

		iLog := letterLogs[i][ii+1:] // all letter logs at i with the exception of the identifier
		jLog := letterLogs[j][ji+1:] // all letter logs at j with the exception of the identifier
		if iLog == jLog {            // This compares identifiers in the case of ties
			// This would be for a situation like the following:
			// letterLogs[i] = "let1 art can"
			// letterLogs[j] = "let3 art can"
			return letterLogs[i] < letterLogs[j]
		}
		return iLog < jLog // This is comparing the logs but without identifiers
	})

	// append both slices, starting with the letter logs first
	return append(letterLogs, digitLogs...)
}

func isDigit(log string) bool {
	b := log[strings.Index(log, " ")+1]
	return '0' <= b && b <= '9'
}
