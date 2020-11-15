package string

import (
	"container/list"
)

// https://leetcode-cn.com/problems/remove-k-digits/
// recursively implementation
func removeKdigits(num string, k int) string {
	l := list.New()
	for i := 0; i < len(num); i++ {
		l.PushBack(num[i])
	}
	_removeKdigits(l, k, 1)

	var min []uint8
	for e := l.Front(); e != nil; e = e.Next() {
		c := (e.Value).(uint8)
		// ascii(48)=0
		// discard leading 0
		if len(min) == 0 && c == 48 {
			continue
		}
		min = append(min, c)
	}
	if len(min) == 0 {
		return "0"
	}
	return string(min)
}

func _removeKdigits(l *list.List, k int, deep int) {
	var next *list.Element
	if l == nil || deep > k {
		return
	}
	for c := l.Front(); c != nil; c = next {
		next = c.Next()
		// number char can compare by ascii.
		if next == nil || (c.Value).(uint8) > (next.Value).(uint8) {
			l.Remove(c)
			break
		}
	}
	_removeKdigits(l, k, deep+1)
}
