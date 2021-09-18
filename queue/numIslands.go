package queue

func NumIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	curIsland := 0

	row := len(grid)
	col := len(grid[0])
	var q MyCircularQueue = ConstructorQueue(row * col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = 1
			}
			if grid[i][j] == '0' {
				grid[i][j] = 0
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				continue
			}
			curIsland++
			bfs(q, grid, i, j)
		}
	}
	return curIsland
}
func bfs(q MyCircularQueue, grid [][]byte, i int, j int) {
	row := len(grid)
	col := len(grid[0])
	grid[i][j] = 0
	q.EnQueue(i*col + j)
	for !q.IsEmpty() {
		cur := q.Front()
		curRow, curCol := cur/col, cur%col
		if curRow > 0 && grid[curRow-1][curCol] == 1 {
			index := (curRow-1)*col + curCol
			q.EnQueue(index)
			grid[curRow-1][curCol] = 0
		}
		if curRow < row-1 && grid[curRow+1][curCol] == 1 {
			index := (curRow+1)*col + curCol
			q.EnQueue(index)
			grid[curRow+1][curCol] = 0
		}
		if curCol > 0 && grid[curRow][curCol-1] == 1 {
			index := curRow*col + curCol - 1
			q.EnQueue(index)
			grid[curRow][curCol-1] = 0
		}
		if curCol < col-1 && grid[curRow][curCol+1] == 1 {
			index := curRow*col + (curCol + 1)
			q.EnQueue(index)
			grid[curRow][curCol+1] = 0
		}
		q.DeQueue()
	}
}
