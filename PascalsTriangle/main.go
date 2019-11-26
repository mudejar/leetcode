package main

func main() {

}

func generate(rowIndex int) [][]int {
	triangle := make([][]int, rowIndex)
	for r := range triangle {
		triangle[r] = make([]int, r+1)
	}
	if rowIndex == 0 {
		return triangle
	}

	triangle[0][0] = 1
	for i := 1; i < rowIndex; i++ {
		helper(triangle, i)
	}

	return triangle
}

func helper(triangle [][]int, row int) {
	prevRowIndex := row-1
	triangle[row][0], triangle[row][row] = 1, 1
	for i := 1; i < len(triangle[row])-1; i++ {
		triangle[row][i] = triangle[prevRowIndex][i-1] + triangle[prevRowIndex][i]
	}
}
