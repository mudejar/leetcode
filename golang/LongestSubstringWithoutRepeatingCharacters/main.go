package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	chars := map[byte]struct{}{}
	maxLength := 0
	j := 0
	i := 0
	length := len(s)
	for i < length && j < length {
		_, ok := chars[s[j]]
		if !ok {
			chars[s[j]] = struct{}{}
			j++
			maxLength = max(maxLength, j-i)
		} else {
			delete(chars, s[i])
			i++
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
