package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		// 使用 go 并发
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		// 接收值
		groupId := <-c
		fmt.Println("Gopher ", groupId, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println(".. sleepyGopher : ", id)
	// 像通道发送值
	c <- id
}
