package main

func main() {

}

func nextPermutation(nums []int) {
	length := len(nums)
	// test if nums[i] < nums[i-1] true
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			swapIndex := i + 1
			for j := i + 2; j < length; j++ {
				if nums[j] > nums[i] {
					swapIndex = j
				} else {
					break
				}
			}
			// swap nums[i], nums[swapIndex]
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			// reverse nums[i+1:]
			for j := i + 1; j < i+1+(length-i-1)/2; j++ {
				nums[j], nums[length-j+i] = nums[length-j+i], nums[j]
			}
			return
		}
	}
	// just reverse it
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
}
