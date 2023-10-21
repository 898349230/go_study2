package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 自定义类型
type person struct {
	Name string
	Age  int
}

// 使用json标签
type person2 struct {
	Name string `json:"nameAttr"`
	Age  int    `json:"ageAttr"`
}

// 自定义类型
type user struct {
	name string
	age  int
}

// 自定义嵌套类型
type admin struct {
	person user
	level  string
}

func main() {
	var bil user
	lisa := user{
		name: "lisa",
		age:  23,
	}
	jenny := user{"Jenny", 18}
	fmt.Println("lisa ", lisa, " bil ", bil, " jenny  ", jenny)

	fred := admin{
		person: user{
			name: "liming",
			age:  19,
		},
		level: "经理",
	}
	fmt.Println("fred : ", fred)

	// 调用方法
	lisa.notify()
	// 也可使用指针调用方法
	lisa1 := &lisa
	lisa1.notify()

	danny := user{"Danny", 22}
	danny.changeAge(18)
	fmt.Println(danny)
	danny2 := &user{"Danny2", 22}
	danny2.changeAge(18)
	fmt.Println(danny2)

	danny3 := user{"Danny3", 22}
	danny3.changeAge1(18)
	// 还是 22
	fmt.Println(danny3)
	danny4 := &user{"danny4", 22}
	danny4.changeAge1(18)
	// 还是22
	fmt.Println(danny4)
	// 打印 struct
	fmt.Printf("danny: %v\n", danny)
	// 带字段值打印
	fmt.Printf("danny: %+v\n", danny)

	// struct 准换为 josn, struct 内的属性应该大写表示可以导出的
	lily := person{
		Name: "lily",
		Age:  18,
	}
	// 使用 json 标签
	lily2 := person2{
		Name: "lily2",
		Age:  19,
	}
	bytes, err := json.Marshal(lily)
	bytes2, err := json.Marshal(lily2)
	exitOnError(err)
	fmt.Println(string(bytes))
	fmt.Println(string(bytes2))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 方法
// 给自定义的类型添加方法
func (u user) notify() {
	fmt.Printf("user notify() -> userName : %s \n", u.name)
}

func (u user) changeAge1(age int) {
	// 因为是复制的值，所以这里修改没有用，不会影响原来的 user
	u.age = age
}

func (u *user) changeAge(age int) {
	u.age = age
}
