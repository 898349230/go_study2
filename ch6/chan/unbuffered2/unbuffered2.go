package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)
	// 为最后一位跑者将计数加1
	wg.Add(1)
	go Runner(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wg.Wait()
}

// 模拟接力比赛中的一位跑着
func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒（阻塞获取通道中的值）
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	// 不是第4位及最后一棒，创建下一位跑者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line \n", runner)
		go Runner(baton)
	}
	// 模拟跑步
	time.Sleep(1000 * time.Millisecond)
	// 第四位跑者  结束比赛
	if runner == 4 {
		fmt.Printf("Runner %d Finished.\n", runner)
		wg.Done()
		return
	}
	// 将接力棒交给下一位跑者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
