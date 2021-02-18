package util

import (
	"math/rand"
	"time"
)

func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	return arr
}

func RandArrayInterface(n int) []interface{} {
	rand.Seed(time.Now().UnixNano())
	var arr []interface{}
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	return arr
}
