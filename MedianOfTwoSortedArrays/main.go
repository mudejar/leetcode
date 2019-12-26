package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"nums1 = [1, 3] nums2 = [2], should result in 2.0":   findMedianSortedArrays([]int{1, 3}, []int{2}),
		"nums1 = [1, 2] nums2 = [3, 4] should result in 2.5": findMedianSortedArrays([]int{1, 2}, []int{3, 4}),
	}).Println()
}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	totalLength := len(nums1) + len(nums2)
//	return findMedianSortedArraysHelper(nums1, nums2, (totalLength-1)/2, (totalLength&1) == 0)
//}

//func findMedianSortedArraysHelper(nums1 []int, nums2 []int, medianIndex int, needNext bool) float64 {
//	// handle length is 0, length can't be both 0
//	if len(nums1) == 0 {
//		medianValue := float64(nums2[medianIndex])
//		if needNext {
//			medianValue += float64(nums2[medianIndex+1])
//			medianValue /= 2.0
//		}
//		return medianValue
//	} else if len(nums2) == 0 {
//		medianValue := float64(nums1[medianIndex])
//		if needNext {
//			medianValue += float64(nums1[medianIndex+1])
//			medianValue /= 2.0
//		}
//		return medianValue
//	}
//
//	// handle medianIndex is 0
//	if medianIndex == 0 {
//		medianValue := float64(0)
//		if nums1[0] == nums2[0] {
//			medianValue = float64(nums1[0])
//		} else if nums1[0] < nums2[0] {
//			medianValue += float64(nums1[0])
//			if needNext {
//				if len(nums1) > 1 {
//					if nums1[1] <= nums2[0] {
//						medianValue += float64(nums1[1])
//					} else {
//						medianValue += float64(nums2[0])
//					}
//				} else {
//					medianValue += float64(nums2[0])
//				}
//				medianValue /= 2.0
//			}
//		} else {
//			medianValue += float64(nums2[0])
//			if needNext {
//				if len(nums2) > 1 {
//					if nums2[1] <= nums1[0] {
//						medianValue += float64(nums2[1])
//					} else {
//						medianValue += float64(nums1[0])
//					}
//				} else {
//					medianValue += float64(nums1[0])
//				}
//				medianValue /= 2.0
//			}
//		}
//		return medianValue
//	}
//
//	// cut the less one from 0 to halfIndex(include), cut anyone if nums1[halfIndex] == nums2[halfIndex]
//	halfIndex := (medianIndex - 1) / 2
//	if len(nums1) < halfIndex+1 {
//		if nums1[len(nums1)-1] <= nums2[halfIndex] {
//			medianIndex -= len(nums1)
//			medianValue := float64(nums2[medianIndex])
//			if needNext {
//				medianValue += float64(nums2[medianIndex+1])
//				medianValue /= 2.0
//			}
//			return medianValue
//		} else {
//			return findMedianSortedArraysHelper(nums1, nums2[halfIndex+1:], medianIndex-halfIndex-1, needNext)
//		}
//	} else if len(nums2) < halfIndex+1 {
//		if nums2[len(nums2)-1] <= nums1[halfIndex] {
//			medianIndex -= len(nums2)
//			medianValue := float64(nums1[medianIndex])
//			if needNext {
//				medianValue += float64(nums1[medianIndex+1])
//				medianValue /= 2.0
//			}
//			return medianValue
//		} else {
//			return findMedianSortedArraysHelper(nums1[halfIndex+1:], nums2, medianIndex-halfIndex-1, needNext)
//		}
//	} else {
//		if nums1[halfIndex] <= nums2[halfIndex] {
//			return findMedianSortedArraysHelper(nums1[halfIndex+1:], nums2, medianIndex-halfIndex-1, needNext)
//		} else {
//			return findMedianSortedArraysHelper(nums1, nums2[halfIndex+1:], medianIndex-halfIndex-1, needNext)
//		}
//	}
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// always search the shorter array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	aLen := len(nums1)
	bLen := len(nums2)
	leftHalfLength := (aLen + bLen + 1) / 2

	// aMinCount and aMaxCount are the min and max number
	// of values A can contribute to the left half of A ∪ B,
	// respectively.
	// Since A is guaranteed to be either shorter or of
	// the same length as B, we know it can contribute
	// 0 or all of its values to the left half of A ∪ B.
	aMinCount, aMaxCount := 0, aLen

	for aMinCount <= aMaxCount {
		// aCount is the number of values A will contribute to left half of A ∪ B
		aCount := aMinCount + ((aMaxCount - aMinCount) / 2)

		// bCount is the number of values B will contribute to the left half of A ∪ B
		// B will contribute as many values as necessary to fill however many remaining
		// slots in the left half.
		bCount := leftHalfLength - aCount

		// Make sure aCount is greater than 0 (because A can contribute 0 values;
		// remember that A is either shorter or of the same length as B). This also
		// implies bCount will be less than B.Length since it won't be possible
		// for B to contribute all of its values if A has contributed at least 1
		// value.
		if aCount > 0 && nums1[aCount-1] > nums2[bCount] { // x > y'
			// Decrease A's contribution size; x lies in the right half of A ∪ B.
			aMaxCount = aCount - 1
		} else if aCount < aLen && nums2[bCount-1] > nums1[aCount] {
			// Make sure aCount is less than A.Length since A can actually contribute
			// all of its values (remember that A is either shorter or of the same
			// length as B). This also implies bCount > 0 because B has to contribute
			// at least 1 value if aCount < A.Length.
			// y > xP

			// Decrease B's contribution size, i.e. increase A's contribution size;
			// y lies in the right half of A ∪ B.
			aMinCount = aCount + 1
		} else {
			//
			// Neither x nor y lie beyond the left half of A ∪ B. This implies we
			// found the right aCount. We don't know how x and y compare to each
			// other yet though.
			//

			var lastValueOnLeft float64
			if aCount == 0 { // all the values come from B
				lastValueOnLeft = float64(nums2[bCount-1])
			} else if bCount == 0 { // all the values come from A
				lastValueOnLeft = float64(nums1[aCount-1])
			} else if nums1[aCount-1] > nums2[bCount-1] {
				lastValueOnLeft = float64(nums1[aCount-1])
			} else {
				lastValueOnLeft = float64(nums2[bCount-1])
			}

			// If aLen + bLen is odd, the median is the greater of x and y.
			if (aLen+bLen)%2 == 1 {
				return lastValueOnLeft
			}

			// If aLen + bLen is even, then the median is the average between
			// the last value contributed on the left side and the first value
			// present on the right half of the array
			var firstValueOnTheRight float64
			if aCount == aLen { // A contributes no values to left half of array
				firstValueOnTheRight = float64(nums2[bCount])
			} else if bCount == bLen {
				firstValueOnTheRight = float64(nums1[aCount])
			} else if nums1[aCount] < nums2[bCount] {
				firstValueOnTheRight = float64(nums1[aCount])
			} else {
				firstValueOnTheRight = float64(nums2[bCount])
			}

			return (firstValueOnTheRight + lastValueOnLeft) / 2
		}
	}

	return -1
}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums1, nums2 = order(nums1, nums2)
//
//	return binarySearch(nums1, nums2, 0, len(nums1)-1)
//}
//
//func binarySearch(numsx []int, numsy []int, start int, end int) float64 {
//	lengthx := len(numsx)
//	lengthy := len(numsy)
//	partitionx := (start + end)/2
//	partitiony := (1 + lengthx + lengthy)/2 - partitionx
//	isEven := (lengthx+lengthy) % 2 == 0
//	log.WithFields(log.Fields{
//		"start": start,
//		"end": end,
//		"partitionx": partitionx,
//		"partitiony": partitiony,
//	}).Println() // debug
//
//	if lengthx == 1 && lengthy == 1 {
//		return (float64(numsx[0])+float64(numsy[0]))/2.0
//	}
//
//	if lengthx == 1 {
//		if numsx[0] <= numsy[partitiony+1] && numsx[0] > numsy[partitiony] {
//			if isEven {
//				return float64(numsx[0])
//			} else {
//				return float64(numsx[0])/float64(numsy[partitiony+1])
//			}
//		} else if numsx[0] > numsy[]
//	}
//
//	if lengthy <= 1 {
//
//	}
//
//	if numsx[partitionx] <= numsy[partitiony+1] && numsy[partitiony] <= numsx[partitionx+1] {
//		return avg(max(numsy[partitiony+1], numsx[partitionx+1]), min(numsy[partitiony], numsx[partitionx]))
//	} else if numsx[partitionx] > numsy[partitiony] {
//		return binarySearch(numsx, numsy, partitionx/2, end) // move left
//	}
//
//	return binarySearch(numsx, numsy, partitionx+1, end) // move right
//}
//
//func max(x, y int) float64 {
//	if x > y {
//		return float64(x)
//	}
//
//	return float64(y)
//}
//
//func min(x, y int) float64 {
//	if x < y {
//		return float64(x)
//	}
//
//	return float64(y)
//}
//
//func avg(x, y float64) float64 {
//	return (x + y)/2
//}
//
//func order(nums1, nums2 []int) ([]int, []int) {
//	if len(nums1) < len(nums2) {
//		return nums1, nums2
//	}
//	return nums2, nums1
//}
