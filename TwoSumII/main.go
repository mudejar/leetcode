package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	table := map[int]int{}

	for i, num := range numbers {
		val, exists := table[num]
		if !exists {
			table[target - num] = i
		} else {
			return []int{val+1, i+1}
		}
	}

	return nil
}
