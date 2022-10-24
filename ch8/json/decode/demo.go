package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个保存键值对的映射
	var c = make(map[string]interface{})
	c["name"] = "zhangsan"
	c["title"] = "我爱你中国"
	c["contact"] = map[string]interface{}{
		"home":   "北京",
		"number": 128900,
	}

	// 将映射序列化到 JSON 字符串, data 是一个 byte 切片
	// MarshalIndent() 生成的json可以设置前缀和缩进
	// data, err := json.MarshalIndent(c, "", "  ")
	// Marshal() 没有缩进
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))
}
