package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// 要反序列化的字符串
var JSON = `{
	 "name": "Gopher",
	 "title": "programmer",
	 "contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
 	}
 }`

func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)

	// json 反序列化到 map 变量
	// interface{} 类型表示可以使用任意类型的值作为给定键的值
	var c2 map[string]interface{}
	err2 := json.Unmarshal([]byte(JSON), &c2)
	if err != nil {
		log.Println("ERROR: ", err2)
		return
	}
	fmt.Println("name: ", c2["Gopher"])
	fmt.Println("title: ", c2["title"])
	fmt.Println("Home: ", c2["contact"].(map[string]interface{})["home"])
	fmt.Println("Cell: ", c2["contact"].(map[string]interface{})["cell"])
}
