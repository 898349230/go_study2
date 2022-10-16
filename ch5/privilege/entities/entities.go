package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	// 嵌套类型是未公开的
	user
	Rights int
}
