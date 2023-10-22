package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	// 创建一个通道, 等待2秒
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		// 2秒内接收到值
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finish sleep")
			// 2秒内没有接收到值
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
