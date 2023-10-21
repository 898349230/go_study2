package main

import "fmt"

func main() {
	// 如果 defer 的函数调用了 recover panic 就会停止，程序继续运行
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("I forgot my towel")
}
