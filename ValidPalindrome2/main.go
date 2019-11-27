package main

func main() {

}

func validPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			if isValidPalindrome(s, start+1, end) {
				return true
			}
			if isValidPalindrome(s, start, end-1) {
				return true
			}
			return false
		}

		start++
		end--
	}

	return true
}

func isValidPalindrome(s string, start int, end int) bool {

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
