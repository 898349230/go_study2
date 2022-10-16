package main

import "fmt"

// 一个通知类行为的接口
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Print("admin notify", a)
}

// notify 使用指针接受者实现的方法
func (u *user) notify() {
	fmt.Println(u)
}

func main() {
	u := user{"张三", "888887@163.com"}
	// 这里会报错，因为 user 类型并没有实现 notifier 接口
	// 如果使用 notify() 的实现使用的是 user ，那么这里可以使用值作为参数传递
	// sendNotification(u)
	sendNotification(&u)

	admin := admin{"李主管", "99999999"}
	sendNotification(&admin)
}

// sendNotification 接受一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
