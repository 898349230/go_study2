package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// feed 模仿期望接收的 XML 文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
 <rss>
	<channel>
		<title>Going Go Programming</title>
		<description>Golang : https://github.com/goinggo</description>
		<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
			<title>Object Oriented Programming Mechanics</title>
			<description>Go is an object oriented language.</description>
			<link>http://www.goinggo.net/2015/03/object-oriented</link>
		</item>
	</channel>
 </rss>`

//  mockServer 模拟互联网上真实服务器的调用
//  mockServer 返回用来处理请求的服务器的指针，
func mockServer() *httptest.Server {
	// f 匿名函数 函数声明符合 http.HandlerFunc 函数类型
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		// 使用 feed 作为响应数据
		fmt.Fprintln(w, feed)
	}
	// HandlerFunc 类型是一个适配器
	// HandlerFunc(f) 是一个处理HTTp请求的Handler对象，内部通过调用f处理请求
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK
	// 使用 mockServer 函数生成模仿服务器
	server := mockServer()
	defer server.Close()
	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode, checkMark)
		}
	}

}
