package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// json 文档
/* {
	"responseData": {
		"results": [
			{
			"GsearchResultClass": "GwebSearch",
			"unescapedUrl": "https://www.reddit.com/r/golang",
			"url": "https://www.reddit.com/r/golang",
			"visibleUrl": "www.reddit.com",
			"cacheUrl": "http://www.google.com/search?q=cache:W...",
			"title": "r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
			"titleNoFormatting": "r/Golang - Reddit",
			"content": "First Open Source \u003cb\u003eGolang\u..."
			},
			{
			"GsearchResultClass": "GwebSearch",
			"unescapedUrl": "http://tour.golang.org/",
			"url": "http://tour.golang.org/",
			"visibleUrl": "tour.golang.org",
			"cacheUrl": "http://www.google.com/search?q=cache:O...",
			"title": "A Tour of Go",
			"titleNoFormatting": "A Tour of Go",
			"content": "Welcome to a tour of the Go programming ..."
			}
		]
	}
} */

type (
	// 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(gr)
}
