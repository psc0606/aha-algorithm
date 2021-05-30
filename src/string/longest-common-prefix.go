package string

// https://leetcode-cn.com/problems/longest-common-prefix/
// wide scan
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// deep scan
func longestCommonPrefixII(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var lcp func(str1, str2 string) string
	lcp = func(str1, str2 string) string {
		if len(str1) > len(str2) {
			return lcp(str2, str1)
		}
		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				return str1[0:i]
			}
		}
		return str1
	}
	ans := lcp(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		ans = lcp(ans, strs[i])
	}
	return ans
}
