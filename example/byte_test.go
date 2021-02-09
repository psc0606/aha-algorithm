package main

import (
	"math"
	"testing"
)

func TestByteMax(t *testing.T) {
	// In addition there two alias types: byte which is the same as uint8 and rune which is the same as int32
	var b byte = 255
	println(b)

	println(math.MaxInt8)
}
