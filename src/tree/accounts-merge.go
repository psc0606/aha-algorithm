package tree

import (
	"sort"
)

// https://leetcode-cn.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	ds := new(DisjointSet)
	n := len(accounts)
	ds.Init(n)

	for i := 0; i < n; i++ {
		ds.MakeSet(i)
	}

	// email -> index
	hashmap := make(map[string]int)
	for i := 0; i < n; i++ {
		// accounts[0] is account name
		for j := 1; j < len(accounts[i]); j++ {
			find, ok := hashmap[accounts[i][j]]
			if ok {
				ds.Union(i, find)
			} else {
				hashmap[accounts[i][j]] = i
			}
		}
	}

	um := make(map[int][]string)
	for k, v := range hashmap {
		root := ds.Find(v)
		accs := um[root]
		if accs == nil {
			accs = []string{accounts[v][0]}
		}
		accs = append(accs, k)
		um[root] = accs
	}

	var result [][]string
	for _, v := range um {
		sort.Strings(v[0:])
		result = append(result, v)
	}
	return result
}
