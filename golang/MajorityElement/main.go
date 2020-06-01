package main

func main() {

}

func majorityElement(nums []int) int {
	mid := len(nums) / 2

	hits := map[int]int{}
	for _, num := range nums {
		hits[num]++
		if hits[num] > mid {
			return num
		}
	}

	return 0
}
