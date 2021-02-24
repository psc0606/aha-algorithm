package backtrace

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// This is a very interesting problem.
// You may think it is very easy, because you can use loop to combine。
// But how many nested levels you will do?
// Actually, backtrace may be a elegant strategy to solved this problem.
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	ht := make(map[byte]string)
	ht['2'] = "abc"
	ht['3'] = "def"
	ht['4'] = "ghi"
	ht['5'] = "jkl"
	ht['6'] = "mno"
	ht['7'] = "pqrs"
	ht['8'] = "tuv"
	ht['9'] = "wxyz"
	return dfsLetter(digits, ht)
}

func dfsLetter(digits string, ht map[byte]string) []string {
	var ret []string
	if len(digits) == 1 {
		alpha := ht[digits[0]]
		for _, c := range alpha {
			ret = append(ret, string(c))
		}
		return ret
	}
	// dfs substring
	arr := dfsLetter(digits[1:], ht)
	alpha := ht[digits[0]]
	for _, c := range alpha {
		for _, str := range arr {
			ret = append(ret, string(c)+str)
		}
	}
	return ret
}
