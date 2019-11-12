package main

import "strconv"

func main() {

}

func fizzBuzz(n int) []string {
	strings := make([]string, 0)
	for i := 1; i <= n; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			strings = append(strings, "FizzBuzz")
		} else if (i % 3 == 0) {
			strings = append(strings, "Fizz")
		} else if (i % 5 == 0) {
			strings = append(strings, "Buzz")
		} else {
			strings = append(strings, strconv.Itoa(i))
		}
	}

	return strings
}