package main

func main() {

}

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		Matrix: matrix,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	this.Matrix[row][col] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := col1; i <= col2; i++ {
		for j := row1; j <= row2; j++ {
			sum += this.Matrix[j][i]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
