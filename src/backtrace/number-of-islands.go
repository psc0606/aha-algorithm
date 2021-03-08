package backtrace

import "aha-algorithm/src/tree"

// https://leetcode-cn.com/problems/number-of-islands/

// Solution1: DFS or BFS
func numIslands(grid [][]byte) int {
	ret := 0
	if grid == nil {
		return ret
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ret++
				dfsNumIslands(grid, m, n, i, j)
			}
		}
	}
	return ret
}

func dfsNumIslands(grid [][]byte, m, n, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0' // never access again, can also use other variable
	dfsNumIslands(grid, m, n, i+1, j)
	dfsNumIslands(grid, m, n, i-1, j)
	dfsNumIslands(grid, m, n, i, j+1)
	dfsNumIslands(grid, m, n, i, j-1)
}

// Solution1: Disjoint set is a great data structure.
// see disjoint-set.go

type Direction struct {
	x, y int
}

func numIslandsByDisjointSet(grid [][]byte) int {
	ret := 0
	if grid == nil {
		return ret
	}
	m, n := len(grid), len(grid[0])
	dirs := []Direction{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	ds := new(tree.DisjointSet)
	ds.Init(m * n)
	for i := 0; i < m*n; i++ {
		ds.MakeSet(i)
	}

	zeros := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// find '1'
			if grid[i][j] == '1' {
				for _, dir := range dirs {
					nextI, nextJ := i+dir.x, j+dir.y
					if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == '1' {
						// store 1-dimension of (i, j)
						ds.Union(i*m+n, nextI*m+nextJ)
					}
				}
				grid[i][j] = '2' // mark as visited
			} else {
				// number of original zeros
				zeros++
			}
		}
	}
	return ds.Count() - zeros
}
