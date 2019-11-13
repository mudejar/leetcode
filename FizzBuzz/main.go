package main

import "strconv"

func main() {

}

//func fizzBuzz(n int) []string {
//	strings := make([]string, 0)
//	for i := 1; i <= n; i++ {
//		if (i % 3 == 0) && (i % 5 == 0) {
//			strings = append(strings, "FizzBuzz")
//		} else if (i % 3 == 0) {
//			strings = append(strings, "Fizz")
//		} else if (i % 5 == 0) {
//			strings = append(strings, "Buzz")
//		} else {
//			strings = append(strings, strconv.Itoa(i))
//		}
//	}
//
//	return strings
//}

// Hashing solution
func fizzBuzz(n int) []string {
	strings := make([]string, 0)
	fizzBuzzDict := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	for i := 1; i <= n; i++ {
		numAndStr := ""
		for _, key := range []int{3, 5} {
			if i % key == 0 {
				numAndStr += fizzBuzzDict[key]
			}
		}

		if numAndStr == "" {
			numAndStr = strconv.Itoa(i)
		}

		strings = append(strings, numAndStr)
	}

	return strings
}