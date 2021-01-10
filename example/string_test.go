package main

import (
	"fmt"
	"strconv"
	"testing"
)

// TODO ??
func TestString(t *testing.T) {
	// 中 's ut8 is 0x4E2D
	str := "ab中"
	byte1 := str[2]    // the first byte of 中 utf8
	byte2 := str[3]    // the second byte of 中 utf8
	byte3 := str[4]    // the third byte of 中 utf8
	fmt.Println(byte1) // 228=0x4E
	fmt.Println(byte2) // 184=0xB8
	fmt.Println(byte3) // 173=0xAD
}

func TestByte2String(t *testing.T) {
	fmt.Println(string(rune(48))) // ascii(48)=0
	fmt.Println(string(rune(97))) // ascii(97)=a
}

func TestStringEqual(t *testing.T) {
	fmt.Println("abc" == "abc") // true
}

func TestIntToString(t *testing.T) {
	// int to string
	str := strconv.Itoa(343)
	fmt.Println(str)
}
