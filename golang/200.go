package leetcode

/*
leetcode 200 : Number of Islands
这道题有点像 word search的简化版.
每遇到'1'后, 开始向四个方向 递归搜索.
搜到后变为'0', 因为相邻的属于一个island. 然后开始继续找下一个'1'.
dfs算法
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				search(grid, i, j)
				count++
			}
		}
	}
	return count
}

func search(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
		return
	}
	grid[x][y] = '0'
	search(grid, x, y-1)
	search(grid, x, y+1)
	search(grid, x-1, y)
	search(grid, x+1, y)
}
