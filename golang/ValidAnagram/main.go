package main

func main() {

}

func isAnagram(s string, t string) bool {
	characters := map[int32]int{}
	for _, char := range s {
		characters[char] += 1
	}

	for _, char := range t {
		characters[char] -= 1
		if characters[char] < 0 {
			return false
		}
	}

	for _, count := range characters {
		if count > 0 {
			return false
		}
	}

	return true
}
