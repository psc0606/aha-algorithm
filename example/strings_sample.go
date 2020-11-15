package main

import (
	"fmt"
	"strings"
)

func main() {
	testTrimLeft()
	testCount()
	testContains()
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
