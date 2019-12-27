package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func closestValue(root *TreeNode, target float64) int {
	if root == nil {
		return math.MaxInt64
	}

	absVal := absValue(target, root.Val)

	rVal := closestValue(root.Right, target)
	absRVal := absValue(target, rVal)

	lVal := closestValue(root.Left, target)
	absLVal := absValue(target, lVal)

	var retValue int
	if absLVal < absRVal {
		retValue = lVal
	} else {
		retValue = rVal
	}

	if absVal < absValue(target, retValue) {
		retValue = root.Val
	}

	return retValue
}

func absValue(target float64, x int) float64 {
	val := target - float64(x)
	if val < 0 {
		return -val
	}
	return val
}
