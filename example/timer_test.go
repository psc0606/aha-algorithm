package main

import (
	"fmt"
	"testing"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 3 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func TestName(t *testing.T) {
	Ping()
}

func Ping() {
	ticker := time.NewTicker(pingPeriod)

	//定义计数器
	count := 1
	fmt.Println("当前时间为:", time.Now(), "count = ", count)
	defer ticker.Stop()
	for {
		// 从定时器中获取数据
		t := <-ticker.C

		fmt.Println("当前时间为:", t, "count = ", count)
		count++
		if count >= 2 {
			break
		}
	}
}
