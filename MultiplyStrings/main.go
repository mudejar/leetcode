package main

func main() {

}

func multiply(num1 string, num2 string) string {
	if len(num1) <= 0 || len(num2) <= 0 || num1 == "0" || num2 == "0" {
		return "0"
	}

	magic := make([]rune, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		v1 := num2[i] - 0x30
		for j := len(num1) - 1; j >= 0; j-- {
			v2 := num1[j] - 0x30
			r := v1 * v2
			magic[len(num2)-i+len(num1)-j-2] += rune(r % 10)
			magic[len(num2)-i+len(num1)-j-1] += rune(r / 10)
		}
	}
	for i, v := range magic {
		magic[i] = v % 10
		v /= 10
		i++
		for v != 0 {
			magic[i] += v % 10
			v /= 10
			i++
		}
	}
	retInRune := make([]rune, 0)
	start := len(num1) + len(num2) - 1
	for magic[start] == 0 {
		start--
	}
	for start >= 0 {
		retInRune = append(retInRune, magic[start]+0x30)
		start--
	}
	return string(retInRune)
}
