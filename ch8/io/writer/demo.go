package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 使用一个 Buffer 值，并将一个字符串写入 Buffer
	// 使用实现 io.Writer 的 Write 方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	fmt.Fprintf(&b, "World!")

	// 将Buffer 的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}
