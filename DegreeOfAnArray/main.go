package main

import "math"

func main() {

}

func findShortestSubArray(nums []int) int {
	degree := 0

	// first create a hashmap that holds slices of indexes which hold each
	// element. Calculate the degree by finding the slice with the most values
	hm := make(map[int][]int)
	for i, v := range nums {
		if _, ok := hm[v]; !ok {
			hm[v] = make([]int, 0)
		}
		hm[v] = append(hm[v], i)

		degree = max(degree, len(hm[v]))
	}

	// now filter our the slices in the previously created hashmap by whichever
	// fit the previously calculated degree, if they do then put them in the
	// lookup slice
	lookup := make([]int, 0)
	for k, v := range hm {
		if len(v) == degree {
			lookup = append(lookup, k)
		}
	}

	distance := math.MaxInt32

	// loop through lookup slice and find the distance between each slice's
	// first and last element. The smallest distance is the answer
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
