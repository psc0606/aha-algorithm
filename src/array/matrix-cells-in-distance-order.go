package array

import "container/list"

// https://leetcode-cn.com/problems/matrix-cells-in-distance-order/submissions/
// calculate distance first, then use distance as array's index
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var arrayList = make([]*list.List, R+C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			// calculate distance first
			distance := abs(r0-i) + abs(c0-j)
			// index into a array by distance
			l := arrayList[distance]
			if l == nil {
				l = list.New()
				arrayList[distance] = l
			}
			l.PushBack([]int{i, j})
		}
	}
	var res [][]int
	for _, l := range arrayList {
		if l == nil {
			continue
		}
		for e := l.Front(); e != nil; e = e.Next() {
			res = append(res, (e.Value).([]int))
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
