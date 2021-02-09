package optimalization

import "math"

// https://leetcode-cn.com/problems/task-scheduler/
// https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-by-leetcode-solution-ur9w/
// A little bit hard to understand.

// Solution one:
// Time: O(∑*n), n is the length of tasks, ∑ is the different number of tasks.
// Space: O(∑)
func leastInterval(tasks []byte, n int) (minTime int) {
	// count every task frequency.
	cnt := make(map[byte]int)
	for _, t := range tasks {
		cnt[t]++
	}

	nextValid := make([]int, 0, len(cnt))
	rest := make([]int, 0, len(cnt))
	for _, c := range cnt {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}

	for range tasks {
		minTime++

		// find the min nextValid[x], if the minTime < nextValid[x], then skip
		minNextValid := math.MaxInt64
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		// skip the idle cpu time
		if minNextValid > minTime {
			minTime = minNextValid
		}

		// find the best nextValid(max nextValid one)
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || rest[i] > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--
	}
	return minTime
}
