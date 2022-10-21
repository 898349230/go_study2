package main

import (
	"log"
	"study2/ch7/work"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

// namePrinter 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个 goroutine 来创建工作池
	pool := work.New(2)
	var wg sync.WaitGroup
	// 切片里的每个名字都会有100goroutine来提交任务
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行，当 Run 返回时我们就知道任务已经处理完成
				pool.Run(&np)
				wg.Done()
			}()
		}
	}

	// 等待所有创建的 goroutine提交他们的工作
	wg.Wait()
	// 让工作池停止工作，等待所有现有的工作完成
	pool.Shutdown()
}
