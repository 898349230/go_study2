package main

import "fmt"

func main() {

	// 创建切片 长度和容量都是5的切片
	slice := make([]string, 5)
	// 长度是3容量是5的切片， 底层数组的长度是指定的容量, make 预分配
	slice2 := make([]string, 3, 5)
	// 字面量声明切片 长度和容量都是4
	slice3 := []int{1, 2, 3, 4}
	// 使用索引声明切片, 使用空字符串初始化第100个元素，如果 [] 中指定了一个值，那就就是数组
	slice4 := []string{99: ""}

	// nil 切片
	var slice5 []int
	// 空切片
	slice6 := make([]int, 0)

	fmt.Println("slice : ", slice, " slice2 : ", slice2, " slice3 : ", slice3,
		" slice4 : ", slice4, " slice5 : ", slice5, " slice6 : ", slice6)

	// slice7 和 slice8 底层用的同一个数组, slice 切分的索引是字节，不是字符
	slice7 := []int{1, 2, 3, 4, 5}
	// slice8 长度是 2（3-1），容量是4（len(slice7)-1）
	slice8 := slice7[1:3]
	println(slice7[2])
	slice8[1] = 100
	println(slice7[2])

	// append 容量超出后会新建一个切片返回
	newSlice := append(slice7, 101)
	println("len(newSlice) ", len(newSlice), " newSlice[5] ", newSlice[5])
	println("len(slice7) ", len(slice7), " slice7[4] ", slice7[4])

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 长度为1（3-2），容量为2（4-2）的切片，这样限制容量可以在使用append时创建新的数组，避免对原切片的值修改
	slice9 := source[2:3:4]
	fmt.Println("slice9 cap = ", cap(slice9))
	fmt.Println(slice9[0])

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println(append(s1, s2...))
	fmt.Println("测试迭代")
	// 迭代 range 创建每个元素的副本，而不是直接返回该元素的引用
	for idx, value := range slice7 {
		// 这里 value 的地址总是相同的，因为迭代返回的变量是一个迭代过程中根据切片一次赋值的新变量
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice7[idx])
		if idx == 2 {
			// value 是副本，不会修改 slice7 的值
			value = 101
			// 会修改 slice7 的值
			slice7[idx] = 102
		}
	}
	fmt.Println(slice7[2])

	// for 迭代
	fmt.Println("for 迭代测试")
	for index := 2; index < len(slice7); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice7[index])
	}
	// cap() 返回切片容量
	fmt.Println("cap() : ", cap(slice7))

	// 多维切片
	fmt.Println("多维切片")
	slice11 := [][]int{{1}, {10, 11}}
	fmt.Println(slice11)
	slice11[1] = append(slice11[1], 111)
	fmt.Println(slice11)
}
