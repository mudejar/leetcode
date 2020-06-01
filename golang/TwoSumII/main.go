package main

func main() {

}

//func twoSum(numbers []int, target int) []int {
//	table := map[int]int{}
//
//	for i, num := range numbers {
//		val, exists := table[num]
//		if !exists {
//			table[target-num] = i
//		} else {
//			return []int{val + 1, i + 1}
//		}
//	}
//
//	return nil
//}

// This is the faster alternative solution space = O(1) and time = O(n)
func twoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}
