package array

// https://leetcode-cn.com/problems/spiral-matrix-ii/
// 1  2  3
// 4  5  6
// 7  8  9
// Time: O(n^2)
// Space: O(n^2), the answer matrix is O(n^2)
func generateMatrix(n int) [][]int {
	var matrix [][]int
	if n <= 0 {
		return matrix
	}
	// create 2d array
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	rows, columns := len(matrix), len(matrix[0])
	top, bottom := 0, rows-1
	left, right := 0, columns-1
	k := 1
	for top <= bottom && left <= right {
		// left -> right
		for i := left; i <= right; i++ {
			matrix[top][i] = k
			k++
		}
		// top -> bottom
		for j := top + 1; j <= bottom; j++ {
			matrix[j][right] = k
			k++
		}
		// the case [[1,2,3,4]]
		if top < bottom && left < right {
			for i := right - 1; i >= left; i-- {
				matrix[bottom][i] = k
				k++
			}
			for j := bottom - 1; j >= top+1; j-- {
				matrix[j][left] = k
				k++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return matrix
}
