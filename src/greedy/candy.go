package greedy

// https://leetcode-cn.com/problems/candy/
func candy(ratings []int) (ans int) {
	n := len(ratings)
	candy := make([]int, n)
	candy[0] = 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		} else {
			candy[i] = 1
		}
	}
	ans = candy[n-1]
	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			candy[j] = max(candy[j], candy[j+1]+1)
		}
		// record the result by the way
		ans += candy[j]
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
