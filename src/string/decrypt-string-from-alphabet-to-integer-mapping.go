package string

import "strconv"

// https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// s = "10#11#12"
// "jkab"

// "1326#"
// "acz"
func freqAlphabets(s string) string {
	length := len(s)
	var res []uint8
	for i := 0; i < length; {
		if i+2 >= length {
			for i < length {
				res = append(res, s[i]+48)
				i++
			}
			break
		}
		if s[i+2] == '#' {
			num, _ := strconv.Atoi(s[i : i+2])
			// decrypt
			c := 'a' + uint8(num-1)
			res = append(res, c)
			// skip to '#' next
			i += 3
		} else {
			res = append(res, s[i]+48)
			i++
		}
	}
	return string(res)
}
