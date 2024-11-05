package main

// CountLakes counts the number of lakes in the given grid.
func CountLakes(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	lakes := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' && !visited[i][j] {
				lakes++
				Search(grid, visited, i, j)
			}
		}
	}

	return lakes
}

// isWater checks if the cell at (i, j) is water.
func isWater(grid [][]byte, i, j int) bool {
	if i < 0 || i >= len(grid) {
		return false
	}
	if j < 0 || j >= len(grid[0]) {
		return false
	}
	if grid[i][j] != '.' {
		return false
	}
	return true
}

// Search recursively explores the lake using depth-first search.
func Search(grid [][]byte, visited [][]bool, i, j int) {
	// Check if the cell is out of bounds or not water
	if !isWater(grid, i, j) || visited[i][j] {
		return
	}

	visited[i][j] = true
	// Helper to explore the 8 adjacent cells
	neighbors := [][]int{
		{0, 1},   // Top-left
		{1, 0},   // Top
		{-1, 0},  // Top-right
		{0, -1},  // Left
		{1, 1},   // Right
		{-1, -1}, // Bottom-left
		{1, -1},  // Bottom
		{-1, 1},  // Bottom-right
	}
	for _, dir := range neighbors {
		Search(grid, visited, i+dir[0], j+dir[1])
	}
}
