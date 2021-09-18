package queue

func NumIslandsByQueue(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

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
	var q MyCircularQueue = ConstructorQueue(row * col)
	islands := make(map[int]int)
	curIsland := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			_, ok := islands[i*col+j]
			if ok {
				continue
			}
			if grid[i][j] == 0 {
				continue
			}
			curIsland++
			q.EnQueue(i*col + j)
			islands[i*col+j] = curIsland
			for !q.IsEmpty() {
				cur := q.Front()
				curRow, curCol := cur/col, cur%col
				cell := grid[curRow][curCol]
				if cell == 1 {
					if curRow > 0 && grid[curRow-1][curCol] == 1 {
						index := (curRow-1)*col + curCol
						_, ok := islands[index]
						if !ok {
							q.EnQueue(index)
							islands[index] = curIsland
						}
					}
					if curRow < row-1 && grid[curRow+1][curCol] == 1 {
						index := (curRow+1)*col + curCol
						_, ok := islands[index]
						if !ok {
							q.EnQueue(index)
							islands[index] = curIsland
						}
					}
					if curCol > 0 && grid[curRow][curCol-1] == 1 {
						index := curRow*col + curCol - 1
						_, ok := islands[index]
						if !ok {
							q.EnQueue(index)
							islands[index] = curIsland
						}
					}
					if curCol < col-1 && grid[curRow][curCol+1] == 1 {
						index := curRow*col + (curCol + 1)
						_, ok := islands[index]
						if !ok {
							q.EnQueue(index)
							islands[index] = curIsland
						}
					}
				}
				q.DeQueue()
			}
		}
	}
	return curIsland
}
