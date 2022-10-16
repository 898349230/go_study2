package main

import (
	"fmt"
	"study2/ch5/privilege/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}
	// Admin 中 user 是非公开的， 但是 user中的 Name 和 Email 是公开的，可以直接设置
	a.Email = "www.abc.com"
	a.Name = "张三"
	fmt.Println(a)
}
