package _200_number_of_islands

import "container/list"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		nRow = len(grid)
		nCol = len(grid[0])
		res = 0
	)
	var (
		dirX = []int{-1, 0, 1, 0}
		dirY = []int{0, 1, 0, -1}
	)
	visited := make([][]bool, nRow)
	for i := range visited {
		visited[i] = make([]bool, nCol)
	}

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if grid[i][j] == '0' || visited[i][j] {
				continue
			}
			res++
			q := list.New()
			q.PushBack(i * nCol + j)
			for q.Len() != 0 {
				t := q.Front()
				q.Remove(t)
				for k := 0; k < 4; k++ {
					x := t.Value.(int) / nCol + dirX[k]
					y := t.Value.(int) % nCol + dirY[k]
					if x < 0 || x >= nRow || y < 0 || y >= nCol || grid[x][y] == '0' || visited[x][y] {
						continue
					}
					visited[x][y] = true
					q.PushBack(x * nCol + y)
				}
			}
		}
	}
	return res
}
