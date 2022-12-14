package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// r 是一个响应，r.Body 是 io.Reader
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// 创建文件保存响应结果
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用 MultiWriter，可以同事向文件和标准输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)
	// 读出响应的内容，写到两个目的地
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
