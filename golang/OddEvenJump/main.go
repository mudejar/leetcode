package main

func main() {

}

func oddEvenJumps(A []int) int {
	dp1 := make([]bool, len(A))
	dp2 := make([]bool, len(A))
	dp1[len(A)-1] = true
	dp2[len(A)-1] = true

	for i := len(A) - 2; i >= 0; i-- {
		jodd, jeven := -1, -1
		for j := i + 1; j < len(A); j++ {
			if A[j] >= A[i] {
				if jodd == -1 || A[j] < A[jodd] || A[j] == A[jodd] && j < jodd {
					jodd = j
					dp1[i] = dp2[j]
				}
			}
			if A[j] <= A[i] {
				if jeven == -1 || A[j] > A[jeven] || A[j] == A[jeven] && j < jeven {
					jeven = j
					dp2[i] = dp1[j]
				}
			}
		}
	}

	res := 0
	for _, v := range dp1 {
		if v {
			res++
		}
	}

	return res
}
