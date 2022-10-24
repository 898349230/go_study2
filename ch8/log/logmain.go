package main

import (
	"log"
)

const (
	// a, b, c, d 在每个常量声明时都将 1 按左移位 iota 个位置，并且 iota 每操作一次都会 自增1
	a = 1 << iota
	b
	c
	d
)

// 配置日志参数
func init() {
	// 设置前缀
	log.SetPrefix("TRACE: ")
	//
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	//  输出 1， 2， 4， 8
	println(a, b, c, d)

	// 写到标准日志记录器，  日志记录器是多 goroutine 安全的
	log.Println("message")
	// Fatalln 在调用 Println() 后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")
	// Panicln 在调用 Println() 后会接着调用 Panic()，除非程序执行 recover 函数，否则会导致程序打印调用栈后终止
	log.Panicln("panic mesage")
}
