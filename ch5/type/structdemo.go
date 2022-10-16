package main

import "fmt"

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
