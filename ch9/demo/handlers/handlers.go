package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes 为网络服务设置路由
func Routes() {
	// 将 /sendjson 服务端点与 SendJSON 函数绑定到一起
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 返回一个简单的JSON文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "张三",
		Email: "www.3839@qq.comr",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
