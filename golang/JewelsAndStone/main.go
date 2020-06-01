package main

func main() {

}

func numJewelsInStones(J string, S string) int {
	jewels := map[int32]struct{}{}

	for _, char := range J {
		jewels[char] = struct{}{}
	}

	total := 0
	for _, char := range S {
		_, exists := jewels[char]
		if exists {
			total++
		}
	}

	return total
}
