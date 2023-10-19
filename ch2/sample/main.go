package main

import (
	"log"
	"os"

	// 下划线 为了 Go 语言对包做初始化操作（执行 init 方法），但是并不使用包里的标识符
	_ "study2/ch2/sample/matchers"
	"study2/ch2/sample/search"
)

// init 在main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("whistleblower ")
}
