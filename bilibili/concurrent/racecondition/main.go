package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var mu sync.Mutex

func worker() {
	n := 0
	// 1s 定时channel
	next := time.After(time.Second)
	// for 代替事件循环
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func main() {
	go worker()
	// 互斥锁
	// mu.Lock()
	// defer mu.Unlock()
}
