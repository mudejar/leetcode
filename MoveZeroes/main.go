package main

func main() {

}

func moveZeroes(nums []int) {
	length := len(nums) -1
	zerosCount, nPos := 0, 0

	for _, n := range nums {
		if n == 0 {
			zerosCount += 1
		} else {
			nums[nPos] = n
			nPos += 1
		}
	}

	for i := length; i >= 0 && zerosCount > 0; i-- {
		nums[i] = 0
		zerosCount -= 1
	}
}
