package main

type ArrayReader interface {
	get(int) int
}

func main() {

}

const (
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

	rightBound := 2
	for ; rightBound <= reader.get(rightBound); rightBound = rightBound * rightBound {
		if reader.get(rightBound) == max {
			break
		}
	}

	length := searchForLengthOfArray(reader, 0, rightBound)

	return binarySearch(reader, target, 0, length-1)
}

// TODO: try doing this iteratively instead
func searchForLengthOfArray(reader ArrayReader, left, right int) int {
	if reader.get(left) != max && reader.get(left+1) == max {
		return left + 1
	}

	midpoint := left + (right-left)/2
	if reader.get(midpoint) == max {
		return searchForLengthOfArray(reader, left, midpoint)
	} else {
		return searchForLengthOfArray(reader, midpoint, right)
	}
}

func binarySearch(reader ArrayReader, target, left, right int) int {
	if right-left <= 1 {
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
