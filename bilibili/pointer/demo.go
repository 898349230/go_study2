package main

import "fmt"

func main() {
	answer := 42
	// & 获取变量的内存地址  * 解引用, 放在变量前面表示解引用，提供内存地址指向的值，
	fmt.Println(&answer)
	address := &answer
	fmt.Println(*address)
	fmt.Println(*&answer)

	canada := "Canada"
	// 声明 home 是指向string类型的指针,* 放在类型前面表示声明指针类型
	var home *string
	// %T 输出类型
	fmt.Printf("home is a %T\n", home)
	home = &canada
	fmt.Println(*home)
}
