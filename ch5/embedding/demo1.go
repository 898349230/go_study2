package main

import "fmt"

type user struct {
	name  string
	email string
}

//  定义一个接口
type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Println("user notify()", u)
}

func (ad *admin) notify() {
	fmt.Println("ad notify()", ad)
}

type admin struct {
	// user 嵌入到 admin 中
	user
	level string
}

func main() {
	ad := admin{
		user:  user{"小张", "www.abc.com"},
		level: "总经理",
	}
	// 可以直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法也被提升到外部类型
	ad.notify()
	// 如果 admin 没有显示的实现 notifier接口，
	// 那么 user 内部类型实现的接口自动提升到外部类型，意味着外部类型也同样实现了这个接口
	sendNotification(&ad)
}

// 接收 notifier类型的值 发送通知
func sendNotification(n notifier) {
	fmt.Println("sendNotification()...")
	n.notify()
}
