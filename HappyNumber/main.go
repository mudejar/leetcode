package main

func main() {

}

func isHappy(n int) bool {
	seen := map[int]struct{}{}
	for n != 1 {
		_, exists := seen[n]
		if exists { // check if the set contains the number calculated so far
			return false
		}

		seen[n] = struct{}{}
		n = getNext(n) // get the next number that is calculated for the happy number link
	}
	return n == 1
}

func getNext(n int) int {
	totalSum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSum += d * d
	}

	return totalSum
}
