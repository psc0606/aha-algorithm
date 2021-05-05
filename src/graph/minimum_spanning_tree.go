package graph

import (
	"aha-algorithm/src/tree"
	"sort"
)

/**
 * 一个有n户人家的村庄，有m条路连接着。村里现在要修路，每条路都有一个代价，现在请你帮忙计算下，最少需要花费多少的代价，就能让这n户人家连接起来。
 * 输入: 3,3,[[1,3,3],[1,2,1],[2,3,1]]
 * 输出: 2
 *
 * https://www.nowcoder.com/practice/735a34ff4672498b95660f43b7fcd628
 * 返回最小的花费代价使得这n户人家连接起来
 * @param n int n户人家的村庄
 * @param m int m条路
 * @param cost int二维数组 一维3个参数，表示连接1个村庄到另外1个村庄的花费的代价
 * @return int
 */
func miniSpanningTree(n int, m int, cost [][]int) int {
	ans := 0
	sort.Sort(Int2dSlice(cost))

	// init union find set
	ds := new(tree.DisjointSet)
	ds.Init(n + 1)
	for i := 0; i < n+1; i++ {
		ds.MakeSet(i)
	}

	for i := 0; i < m; i++ {
		edge := cost[i]
		if ds.Same(edge[0], edge[1]) {
			continue
		}
		ans += edge[2]
		ds.Union(edge[0], edge[1])
	}
	return ans
}

type Int2dSlice [][]int

func (s Int2dSlice) Len() int           { return len(s) }
func (s Int2dSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int2dSlice) Less(i, j int) bool { return s[i][2] < s[j][2] }
