package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 通知正在执行的 goroutine停止工作的标志
	shutdown int64
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

// 测试并发导致的问题  竞态
func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)
	// 创建两个 goroutine
	go doWork("A")
	go doWork("B")
	// 给 goroutine 执行的时间
	time.Sleep(1 * time.Second)
	// 等待 gouroutine 结束
	fmt.Println("Shutdown Now")
	// 安全的设置 shutdown 标志
	// goroutine 同时调用 LoadInt64 和 StoreInt64时会同步调用，保证操作是安全的，不会进入竞态状态
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

// doWork 模拟执行工作的 goroutine，检测之前的 shutdown 标志来决定是否提前终止
func doWork(name string) {
	//函数退出时调用Done()来通知main函数工作已经完成
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}

}
