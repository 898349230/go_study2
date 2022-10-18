package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// 要使用的goroutine数量
	numberGoroutines = 4
	// 要处理的工作的数量
	taskLoad = 10
)

var wg sync.WaitGroup

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)
	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	// 当所有工作都处理完时关闭通道，以便所有goroutine退出
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}

// worker 作为goroutine启动来处理，从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		// 获取通道的值
		task, ok := <-tasks
		if !ok {
			// 意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d: Shutting Down \n", worker)
			return
		}
		// 开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		//     完成工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)

	}
}
