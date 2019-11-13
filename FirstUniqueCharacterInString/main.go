package main

func main() {

}

func firstUniqChar(s string) int {
	uniqueChars := map[int32]bool{}
	for _, char := range s {
		_, exists := uniqueChars[char]
		if exists {
			uniqueChars[char] = false
		} else {
			uniqueChars[char] = true
		}
	}

	for i, val := range s {
		if uniqueChars[val] {
			return i
		}
	}

	return -1
}