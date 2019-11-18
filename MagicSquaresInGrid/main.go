package main

func main() {

}

func numMagicSquaresInside(grid [][]int) int {

	count := 0
	row := len(grid)
	col := len(grid[0])

	for i := 1; i < row - 1; i++ {
		for j := 1; j < col - 1; j++ {
			A, B, C := grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1]
			D, E, F := grid[i][j-1], grid[i][j], grid[i][j+1]
			G, H, I := grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1]
			distinct := true
			tmp := []int{A, B, C, D, E, F, G, H, I}
			temp := make(map[int]int)
			for _, v := range tmp {
				temp[v]++
			}
			for i, v := range temp {
				if i > 9 || i < 1 || v > 1 {
					distinct = false
				}
			}
			a := grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1]
			b := grid[i][j-1] + grid[i][j] + grid[i][j+1]
			c := grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1]

			d := grid[i-1][j-1] + grid[i][j-1] + grid[i+1][j-1]
			e := grid[i-1][j] + grid[i][j] + grid[i+1][j]
			f := grid[i-1][j+1] + grid[i][j+1] + grid[i+1][j+1]

			g := grid[i][j] + grid[i-1][j-1] + grid[i+1][j+1]
			h := grid[i][j] + grid[i-1][j+1] + grid[i+1][j-1]
			if a == b && b == c && c == d && d == e && e == f && f == g && g == h && distinct {
				count++
			}
		}
	}
	return count
}