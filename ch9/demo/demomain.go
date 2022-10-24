package main

import (
	"log"
	"net/http"
	"study2/ch9/demo/handlers"
)

func main() {
	handlers.Routes()
	log.Println("listener : Started : Listening on :4000")
	// 显示服务监控的端口，并且启动网络服务，等待请求
	// 本地浏览器可以访问 http://localhost:4000/sendjson
	http.ListenAndServe(":4000", nil)
}
