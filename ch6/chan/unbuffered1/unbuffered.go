package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)
	// 计数加2， 表示要等待两个 goroutine
	wg.Add(2)
	// 启动两个选手
	go player("Jenny", court)
	go player("Liming", court)

	//发球
	court <- 1
	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		// 等待球被击打过来（从通道中获取值）
		// 这里会阻塞直到从管道中获取值
		ball, ok := <-court
		if !ok {
			// 如果通道关闭，就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手（通道内放入值）
		// 这里两个goroutine都会阻塞，直到交换数据完成
		court <- ball
	}
}
