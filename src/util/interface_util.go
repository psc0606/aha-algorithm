package util

import "reflect"

// If a interface is nil, we can not use a == nil, it will give `false`.
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
