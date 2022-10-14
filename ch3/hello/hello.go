// 当编译器发现某个包的名字为 main 时，也一定也会发现名为 main()的Ѧ数，否则不会创建可执行文件
package main

import (
	// 使用 ffmt 别名, 可以解决报名冲突问题
	"code/chapter3/words"
	ffmt "fmt"
	"io/ioutil"
	"os"
)

// main 函数执行前执行
func init() {
	// sql.Register("", new (PostgresDriver))
	ffmt.Println("init()")
}

func main() {
	ffmt.Println("Hello World")
	fileName := os.Args[1]
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		ffmt.Println(err)
		return
	}
	text := string(contents)
	count := words.CountWords(text)
	ffmt.Println("There are %d words in your text", count)
}
