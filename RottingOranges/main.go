package main

func main() {

}

type Node struct {
	Val        int
	Row        int
	Column     int
	TreeHeight int
	Children   []*Node
}

type SearchTree struct {
	Grid             [][]int
	HasRottenOranges bool
	Height           int
	visited          map[string]struct{}
	children         []*Node
}

func NewTree(grid [][]int, node *Node) *SearchTree {
	return &SearchTree{
		Grid:             grid,
		HasRottenOranges: false,
		Height:           0,
		visited:          make(map[string]struct{}),
		children:         []*Node{node},
	}
}

func (t *SearchTree) NextBreadthSearchLevel() {
	totalChildren := len(t.children)
	i := 0
	for ; i < totalChildren; i++ {
		top := t.children[i]
		t.visited[top.Serialize()] = struct{}{}
		t.Grid[top.Row][top.Column] = 2
		top.loadChildren(t.Grid, t.visited)
		if top.TreeHeight > t.Height {
			t.Height = top.TreeHeight
		}
		t.children = append(t.children, top.Children...)
	}
	t.children = t.children[i:]
}

func (t *SearchTree) IsFullyFormed() bool {
	if t.children == nil {
		return true
	}

	if len(t.children) == 0 {
		return true
	}

	return false
}

func allOrangesAreRotten(grid [][]int) bool {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == 1 {
				return false
			}
		}
	}

	return true
}

func (n *Node) Serialize() string {
	return string(n.Row) + string(n.Column)
}

func NewNode(val int, r int, c int, h int) *Node {
	return &Node{
		Val:        val,
		Row:        r,
		Column:     c,
		TreeHeight: h,
	}
}

func (node *Node) loadChildren(grid [][]int, visited map[string]struct{}) *Node {
	w := node.Column - 1
	e := node.Column + 1
	n := node.Row + 1
	s := node.Row - 1

	if w >= 0 && grid[node.Row][w] == 1 {
		node.addChild(grid, node.Row, w, node.TreeHeight+1, visited)
	}

	if e < len(grid[0]) && grid[node.Row][e] == 1 {
		node.addChild(grid, node.Row, e, node.TreeHeight+1, visited)
	}

	if n < len(grid) && grid[n][node.Column] == 1 {
		node.addChild(grid, n, node.Column, node.TreeHeight+1, visited)
	}

	if s >= 0 && grid[s][node.Column] == 1 {
		node.addChild(grid, s, node.Column, node.TreeHeight+1, visited)
	}

	return node
}

func (n *Node) addChild(grid [][]int,
	row int,
	column int,
	height int,
	visited map[string]struct{},
) {
	child := NewNode(grid[row][column], row, column, height)
	_, ok := visited[child.Serialize()]
	if !ok {
		n.Children = append(n.Children, child)
		grid[row][column] = 2
	}
}

func orangesRotting(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	if rowLen == 1 && colLen == 1 {
		if grid[0][0] == 0 {
			return 0
		}

		if grid[0][0] == 1 {
			return -1
		}
	}

	// first find every rotten orange and use each as a head node for a search
	// tree
	allRottenOranges := []*SearchTree{}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == 2 {
				allRottenOranges = append(allRottenOranges,
					NewTree(
						grid,
						NewNode(grid[i][j], i, j, 0),
					),
				)
			}
		}
	}

	// A spoiled batch is a collection of rotten oranges that can't spread any
	// further. This is equivalent to a fully formed search tree
	spoiledBatches := 0
	for spoiledBatches < len(allRottenOranges) {
		spoiledBatches = 0
		for _, rottenOrangeSearchTree := range allRottenOranges {
			if rottenOrangeSearchTree != nil {
				if !rottenOrangeSearchTree.IsFullyFormed() {
					rottenOrangeSearchTree.NextBreadthSearchLevel()
				} else {
					spoiledBatches++
				}
			}
		}
	}

	// check if any leftover oranges have been isolated and didn't rot
	if !allOrangesAreRotten(grid) {
		return -1
	}

	// calculate the total steps by comparing the height of the fully formed
	// search trees
	totalSteps := 0
	for _, orange := range allRottenOranges {
		if totalSteps < orange.Height {
			totalSteps = orange.Height
		}
	}

	return totalSteps
}
