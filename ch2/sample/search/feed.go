package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	// `json:site` 称作标记，描述了JSON解码的原数据，每个标及将结构类型里字段对应到 JSON 文档里指定名字的字段。
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 关闭文件
	defer file.Close()
	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
