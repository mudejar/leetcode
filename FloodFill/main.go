package main

func main() {

}

var (
	R int
	C int
	color int
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	R, C = len(image), len(image[0])
	color = image[sr][sc]
	if color == newColor {
		return image
	}

	dfs(image, newColor, sr, sc)
	return image
}

func dfs(image [][]int, newColor int, r int, c int) {
	if image[r][c] == color {
		image[r][c] = newColor
		if r >= 1 {
			dfs(image, newColor, r-1, c)
		}
		if r+1 < R {
			dfs(image, newColor, r+1, c)
		}
		if c >= 1 {
			dfs(image, newColor, r, c-1)
		}
		if c + 1 < C {
			dfs(image, newColor, r, c+1)
		}
	}
}