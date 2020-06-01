package main

import "strconv"

func main() {

}

func addStrings(num1 string, num2 string) string {
	num1Length := len(num1)
	num2Length := len(num2)
	zeros := ""
	if num1Length > num2Length {
		for i := 0; i < num1Length-num2Length; i++ {
			zeros += "0"
		}
		num2 = zeros + num2
	} else {
		for i := 0; i < num2Length-num1Length; i++ {
			zeros += "0"
		}
		num1 = zeros + num1
	}

	str := ""
	carry := 0
	sum := 0
	for i := len(num1) - 1; i >= 0; i-- {
		one, _ := strconv.Atoi(string(num1[i]))
		two, _ := strconv.Atoi(string(num2[i]))
		sum = one + two + carry
		carry = sum / 10
		sum = sum % 10
		str = strconv.Itoa(sum) + str
	}

	if carry != 0 {
		str = strconv.Itoa(carry) + str
	}

	return str
}
