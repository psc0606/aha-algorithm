package array

import "sort"

// https://leetcode-cn.com/problems/group-anagrams/
// Sort + hashtable
func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	ht := make(map[string][]string)
	for _, s := range strs {
		chars := []byte(s)
		sort.Sort(ByteSlice(chars))
		v, ok := ht[string(chars)]
		if ok {
			v = append(v, s)
			ht[string(chars)] = v
		} else {
			ht[string(chars)] = []string{s}
		}
	}
	for _, v := range ht {
		ret = append(ret, v)
	}
	return ret
}

type ByteSlice []byte

func (s ByteSlice) Len() int           { return len(s) }
func (s ByteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByteSlice) Less(i, j int) bool { return s[i] < s[j] }
