package main

import (
	"fmt"
)

func main() {
	// 声明映射
	dick := make(map[string]int)
	// 预留 8 个空间
	dick2 := make(map[string]int, 8)
	dick2["1"] = 1
	colors := map[string]string{"red": "#da1337", "Orange": "#e95a22"}
	dick3 := map[string][]int{}
	fmt.Println(dick, colors, dick3)
	// 设置值
	dick["red"] = 1
	fmt.Println(dick)
	// nil映射
	var m1 map[string]int
	fmt.Println("m1 : ", m1)
	// 获取值
	val, exist := colors["red"]
	if exist {
		fmt.Printf("val : %s", val)
	}

	if val, ok := colors["red"]; ok {
		fmt.Printf("valred : %s", val)
	} else {

	}

	// 迭代
	for key, val := range colors {
		fmt.Printf("key : %s , value : %s", key, val)
	}
	// 删除键值对
	fmt.Println()
	delete(colors, "red")
	for key, val := range colors {
		fmt.Printf("key : %s , value : %s", key, val)
	}
}
