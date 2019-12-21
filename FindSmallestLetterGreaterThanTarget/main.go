package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField(`letters = ["c","f","j"] target = "d"`, string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))).Println()
}

func nextGreatestLetter(letters []byte, target byte) byte {
	if target < letters[0] {
		return letters[0]
	}

	if target >= letters[len(letters)-1] {
		return letters[0]
	}

	letters = removeDuplicates(letters)

	return search(letters, target, 0, len(letters)-1)
}

func removeDuplicates(letters []byte) []byte {
	arr := []byte{}
	charTable := map[byte]struct{}{}

	for _, letter := range letters {
		_, exists := charTable[letter]
		if !exists {
			arr = append(arr, letter)
			charTable[letter] = struct{}{}
		}
	}

	return arr
}

func search(letters []byte, target byte, left int, right int) byte {
	if left > right {
		return letters[left]
	}

	midpoint := (left + right) / 2
	if letters[midpoint] < target {
		return search(letters, target, midpoint+1, right)
	}

	if letters[midpoint] > target {
		return search(letters, target, left, midpoint-1)
	}

	return letters[midpoint+1]
}
