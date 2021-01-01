package optimalization

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// First: dynamic programming
// Time: O(n)
// Space: O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	hashmap := make(map[uint8]int)
	// init the first one
	hashmap[s[0]] = 0
	startIndexOfSubstring := 0
	currMaxLength := 1

	// start from second
	for i := 1; i < len(s); i++ {
		index, ok := hashmap[s[i]]
		if ok {
			// `index` may be out of date.
			startIndexOfSubstring = util.Max(startIndexOfSubstring, index+1)
		}
		currMaxLength = util.Max(i-startIndexOfSubstring+1, currMaxLength)
		hashmap[s[i]] = i
	}
	return currMaxLength
}

// Second: Another perfect way to solve this problem is sliding window.
// Time: O(n)
// Space: O(n)
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	hashmap := make(map[uint8]int)
	// The max size of window
	maxWindowSize := 0
	// The start index of the current window.
	j := 0
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		// more than one in current window, shrink the left until no more duplicate elements in the current window.
		for hashmap[s[i]] > 1 {
			hashmap[s[j]]--
			if hashmap[s[j]] <= 0 {
				delete(hashmap, s[j])
			}
			j++
		}
		maxWindowSize = util.Max(maxWindowSize, i-j+1)
	}
	return maxWindowSize
}
