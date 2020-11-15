package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	// str1 == str2, str1 > str2, str1 < str2
	fmt.Println(strings.Compare("abc", "abd"))
}

func TestTrimLeft(t *testing.T) {
	fmt.Println(strings.TrimLeft("001200", "0"))
}

func TestCount(t *testing.T) {
	fmt.Println(strings.Count("001200", "0"))
}

func TestContains(t *testing.T) {
	fmt.Println(strings.Contains("001200", "00"))
}
