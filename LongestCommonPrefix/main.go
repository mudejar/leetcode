package main

import (
	log "github.com/sirupsen/logrus"
	"math"
)

func main() {
	log.WithFields(log.Fields{
		`["flower","flow","flight"]`: longestCommonPrefix([]string{"flower","flow","flight"}),
		`["dog","racecar","car"]`: longestCommonPrefix([]string{"dog","racecar","car"}),
	}).Println()
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	commonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		smallest := int(math.Min(float64(len(str)), float64(len(commonPrefix))))
		commonPrefix = commonPrefix[:smallest]
		for j := 0; j < smallest; j++ {
			if commonPrefix[j] != str[j] {
				commonPrefix = commonPrefix[:j]
				break
			}
		}
	}
	return commonPrefix
}