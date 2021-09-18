package queue

func NumIslandsDfs(grid [][]byte) int {
	curIsland := 0

	row := len(grid)
	col := len(grid[0])

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		if grid[i][j] == '1' {
	// 			grid[i][j] = 1
	// 		}
	// 		if grid[i][j] == '0' {
	// 			grid[i][j] = 0
	// 		}
	// 	}
	// }
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				continue
			}
			curIsland++
			dfs(grid, i, j)
		}
	}
	return curIsland
}
func dfs(grid [][]byte, i int, j int) {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || i > row-1 || j < 0 || j > col-1 || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
