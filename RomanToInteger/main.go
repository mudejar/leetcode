package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"III=3":        romanToInt("III"),
		"IV=4":         romanToInt("IV"),
		"IX=9":         romanToInt("IX"),
		"LVIII=58":     romanToInt("LVIII"),
		"MCMXCIV=1994": romanToInt("MCMXCIV"),
	}).Println()
}

func romanToInt(s string) int {
	var arabic = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	pairs := map[string]struct{}{
		"IV": struct{}{},
		"IX": struct{}{},
		"XL": struct{}{},
		"XC": struct{}{},
		"CD": struct{}{},
		"CM": struct{}{},
	}

	result := 0
	prev := ""
	for i := 0; i < len(s); i++ {
		result += arabic[string(s[i])]
		_, exists := pairs[prev+string(s[i])]
		if exists {
			result -= 2 * arabic[prev]
		}
		prev = string(s[i])
	}
	return result
}
