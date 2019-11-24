package main

import(
	"regexp"
	"strings"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("MCW", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})).Println()
}

func mostCommonWord(paragraph string, banned []string) string {
	words := getWords(paragraph)

	bannedList := map[string]struct{}{}
	for _, word := range banned {
		bannedList[word] = struct{}{}
	}

	mostCommonWord := ""
	max := 0
	instances := map[string]int{}
	for _, word := range words {
		_, exists := bannedList[word]
		if !exists {
			instances[word]++
			if instances[word] > max {
				max = instances[word]
				mostCommonWord = word
			}
		}
	}

	return mostCommonWord
}


func getWords(paragraph string) []string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	paragraph = reg.ReplaceAllString(paragraph, ",")
	paragraph = strings.ToLower(paragraph)
	return strings.Split(paragraph, ",")
}