// 当编译器发现某个包的名字为 main 时，也一定也会发现名为 main()的Ѧ数，否则不会创建可执行文件
package main

import (
	// 使用 ffmt 别名, 可以解决报名冲突问题
	"code/chapter3/words"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fileName := os.Args[1]
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(contents)
	count := words.CountWords(text)
	fmt.Println("There are %d words in your text \n", count)
}
