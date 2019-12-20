package main

type ArrayReader interface {
	get(int) int
}

func main() {

}

const(
	max = 2147483647
)

func search(reader ArrayReader, target int) int {
	if reader.get(0) == max {
		return -1
	}

	if reader.get(1) == max {
		if reader.get(0) == target {
			return 0
		}
		return -1
	}

	length := 10000
	for ;; length-- {
		if reader.get(length) == max {break}
	}

	return binarySearch(reader, target, 0, length)
}

func binarySearch(reader ArrayReader, target, left, right int) int {
	if right - left <= 1 {
		if reader.get(left) == target {
			return left
		}
		if reader.get(right) == target {
			return right
		}
		return -1
	}
	midpoint := (left + right) / 2
	if reader.get(midpoint) == target {
		return midpoint
	}

	if reader.get(midpoint) > target {
		return binarySearch(reader, target, left, midpoint-1)
	}

	return binarySearch(reader, target, midpoint+1, right)
}