package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "bad apple", "goodbye all"} {
		downstream <- v
	}

	// downstream <- ""
	// 关闭通道
	close(downstream)
}

// func filterGopher(upstream, downstream chan string) {
// 	for {
// 		// item := <-upstream
// 		// if item == "" {
// 		// 	downstream <- ""
// 		// 	return
// 		// }

// 		// 判断通道是否关闭
// 		item, ok := <-upstream
// 		if !ok {
// 			// 如果上游通道关闭，就关闭下游通道
// 			close(downstream)
// 			return
// 		}
// 		if !strings.Contains(item, "bad") {
// 			downstream <- item
// 		}
// 	}
// }

func filterGopher(upstream, downstream chan string) {
	// 使用 for 循环， 直到通道关闭
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
