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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	return findMedianSortedArraysHelper(nums1, nums2, (totalLength-1)/2, (totalLength&1) == 0)
}

func findMedianSortedArraysHelper(nums1 []int, nums2 []int, medianIndex int, needNext bool) float64 {
	// handle length is 0, length can't be both 0
	if len(nums1) == 0 {
		medianValue := float64(nums2[medianIndex])
		if needNext {
			medianValue += float64(nums2[medianIndex+1])
			medianValue /= 2.0
		}
		return medianValue
	} else if len(nums2) == 0 {
		medianValue := float64(nums1[medianIndex])
		if needNext {
			medianValue += float64(nums1[medianIndex+1])
			medianValue /= 2.0
		}
		return medianValue
	}

	// handle medianIndex is 0
	if medianIndex == 0 {
		medianValue := float64(0)
		if nums1[0] == nums2[0] {
			medianValue = float64(nums1[0])
		} else if nums1[0] < nums2[0] {
			medianValue += float64(nums1[0])
			if needNext {
				if len(nums1) > 1 {
					if nums1[1] <= nums2[0] {
						medianValue += float64(nums1[1])
					} else {
						medianValue += float64(nums2[0])
					}
				} else {
					medianValue += float64(nums2[0])
				}
				medianValue /= 2.0
			}
		} else {
			medianValue += float64(nums2[0])
			if needNext {
				if len(nums2) > 1 {
					if nums2[1] <= nums1[0] {
						medianValue += float64(nums2[1])
					} else {
						medianValue += float64(nums1[0])
					}
				} else {
					medianValue += float64(nums1[0])
				}
				medianValue /= 2.0
			}
		}
		return medianValue
	}

	// cut the less one from 0 to halfIndex(include), cut anyone if nums1[halfIndex] == nums2[halfIndex]
	halfIndex := (medianIndex - 1) / 2
	if len(nums1) < halfIndex+1 {
		if nums1[len(nums1)-1] <= nums2[halfIndex] {
			medianIndex -= len(nums1)
			medianValue := float64(nums2[medianIndex])
			if needNext {
				medianValue += float64(nums2[medianIndex+1])
				medianValue /= 2.0
			}
			return medianValue
		} else {
			return findMedianSortedArraysHelper(nums1, nums2[halfIndex+1:], medianIndex-halfIndex-1, needNext)
		}
	} else if len(nums2) < halfIndex+1 {
		if nums2[len(nums2)-1] <= nums1[halfIndex] {
			medianIndex -= len(nums2)
			medianValue := float64(nums1[medianIndex])
			if needNext {
				medianValue += float64(nums1[medianIndex+1])
				medianValue /= 2.0
			}
			return medianValue
		} else {
			return findMedianSortedArraysHelper(nums1[halfIndex+1:], nums2, medianIndex-halfIndex-1, needNext)
		}
	} else {
		if nums1[halfIndex] <= nums2[halfIndex] {
			return findMedianSortedArraysHelper(nums1[halfIndex+1:], nums2, medianIndex-halfIndex-1, needNext)
		} else {
			return findMedianSortedArraysHelper(nums1, nums2[halfIndex+1:], medianIndex-halfIndex-1, needNext)
		}
	}
}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums1, nums2 = order(nums1, nums2)
//
//	return binarySearch(nums1, nums2, 0, len(nums1)-1)
//}

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
