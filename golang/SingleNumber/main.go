package main

func main() {

}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := 0
	for _, n := range nums {
		result = result ^ n
	}

	return result
}
