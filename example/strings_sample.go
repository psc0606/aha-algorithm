package main

import (
	"fmt"
	"strings"
)

func main() {
	testCompare()
	testTrimLeft()
	testCount()
	testContains()
}

func testCompare() {
	// str1 == str2, str1 > str2, str1 < str2
	fmt.Println(strings.Compare("abc", "abd"))
}

func testTrimLeft() {
	fmt.Println(strings.TrimLeft("001200", "0"))
}

func testCount() {
	fmt.Println(strings.Count("001200", "0"))
}

func testContains() {
	fmt.Println(strings.Contains("001200", "00"))
}
