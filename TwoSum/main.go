package main

func main() {

}

func twoSum(nums []int, target int) []int {
	remainders := map[int]int{}
	for i, num := range nums {
		index, ok := remainders[target-num]
		if ok {
			return []int{index, i}
		}
		remainders[target-num] = i
	}

	return []int{0, 0}
}
