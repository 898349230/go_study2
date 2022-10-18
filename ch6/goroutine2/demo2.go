package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 每个可用的核心分配一个逻辑处理器
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create Goroutine ")
	go prinPrime("A")
	go prinPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

func prinPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
