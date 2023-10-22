package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// 使用 go 并发
		go sleepyGoHelper(i)
	}

	time.Sleep(4 * time.Second)
}

func sleepyGoHelper(i int) {
	time.Sleep(3 * time.Second)
	fmt.Println(".. sleepyGoHelper : ", i)
}
