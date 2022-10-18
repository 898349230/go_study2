package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int64
	// wg 用来等待程序结束
	wg sync.WaitGroup
)

// 测试并发导致的问题  竞态
func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)
	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)
	// 等待 gouroutine 结束
	wg.Wait()
	// 最后 counter 值是 2
	fmt.Println("Final counter : ", counter)
}

func incCounter(id int) {
	//函数退出时调用Done()来通知main函数工作已经完成
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同步加法，原子性的方法，同一时刻只能有一个 gooutine允许并完成这个加法操作
		atomic.AddInt64(&counter, 1)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}

}
