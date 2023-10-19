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

	arr4 := [...]int{6, 9, 0, 1}
	fmt.Println(arr4)

	for i := 0; i < len(arr4); i++ {
		fmt.Println(i, " ", arr4[i])
	}

	for i, v := range arr4 {
		fmt.Println(i, " ", v)
	}

}
