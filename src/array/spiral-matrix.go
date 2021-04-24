package array

// https://leetcode-cn.com/problems/spiral-matrix/

// 1,2,3
// 4,5,6
// 7,8,9
//
// 1, 2, 3, 4
// 5, 6, 7, 8
// 9,10, 11, 12
// four pointer
func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	var ans []int
	rows, columns := len(matrix), len(matrix[0])
	top, bottom := 0, rows-1
	left, right := 0, columns-1
	for top <= bottom && left <= right {
		// left -> right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		// top -> bottom
		for j := top + 1; j <= bottom; j++ {
			ans = append(ans, matrix[j][right])
		}
		// the case [[1,2,3,4]]
		if top < bottom && left < right {
			for i := right - 1; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			for j := bottom - 1; j >= top+1; j-- {
				ans = append(ans, matrix[j][left])
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return ans
}
