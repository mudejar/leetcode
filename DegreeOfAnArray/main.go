package main

import "math"

func main() {

}

func findShortestSubArray(nums []int) int {
	degree := 0
	hm := make(map[int][]int)
	for i, v := range nums {
		if _, ok := hm[v]; !ok {
			hm[v] = make([]int, 0)
		}
		hm[v] = append(hm[v], i)

		degree = max(degree, len(hm[v]))
	}

	lookup := make([]int, 0)
	for k, v := range hm {
		if len(v) == degree {
			lookup = append(lookup, k)
		}
	}

	distance := math.MaxInt32

	for _, n := range lookup {
		v, _ := hm[n]
		distance = min(distance, (v[len(v)-1]-v[0])+1)
	}

	return distance
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
