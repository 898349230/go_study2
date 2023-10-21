package main

import "fmt"

func proverbs(name string) error {

}

func main() {
	// 声明一个 nil
	var nowhere *int
	fmt.Println(nowhere)
	// 解引用 nil 会 panic
	fmt.Println(*nowhere)

}
