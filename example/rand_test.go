package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
}
