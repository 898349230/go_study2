package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	arr[0] = 10
	fmt.Println(arr)
	// 指针数组
	arr2 := []*int{new(int), new(int)}
	*arr2[0] = 1
	*arr2[1] = 2
	fmt.Println(arr2)
	// 多维数组
	arr3 := [3][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(arr3)

}
