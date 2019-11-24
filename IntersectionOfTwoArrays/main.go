package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	nums1, nums2 = orderByLength(nums1, nums2)

	letters := map[int]struct{}{}
	for _, val := range nums1 {
		letters[val] = struct{}{}
	}

	intersection := []int{}
	for _, val := range nums2 {
		_, exists := letters[val]
		if exists {
			intersection = append(intersection, val)
			delete(letters, val)
		}
	}

	return intersection
}

func orderByLength(n, m []int) ([]int, []int) {
	if len(n) > len(m) {
		return n, m
	}

	return m, n
}