package main

func numIslands(grid [][]byte) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid)
				res++
			}
		}
	}

	return res
}

func inBounds(i, j int, grid [][]byte) bool {
	return (0 <= i && i < len(grid) &&
		0 <= j && j < len(grid[i]))
}

func dfs(i, j int, grid [][]byte) {
	if !inBounds(i, j, grid) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'

	dfs(i-1, j, grid)
	dfs(i, j-1, grid)
	dfs(i+1, j, grid)
	dfs(i, j+1, grid)
}
