package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器
	// runtime.GOMAXPROCS(1)
	// 如果分配两个逻辑处理器，会让goroutine并行运行
	runtime.GOMAXPROCS(2)
	// wg 用来等待程序完成，计数加2， 表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines ")
	// 声明一个匿名函数，并创建一个 goroutine
	// 通过 go 关键字创建 goroutine 执行
	go func() {
		// 函数退出时调用Done函数来通知main函数已经完成
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	//等待 goroutine 执行结束
	wg.Wait()
	fmt.Println("\nTerminationg Program")
}
