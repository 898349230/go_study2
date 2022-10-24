// 包名使用 _test 结尾，测试代码只能访问包里公开的标识符，
// 即使测试代码文件和被测试的代码放在同一个文件夹中，也只能访问公开的标识符
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"study2/ch9/demo/handlers"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	// 为服务端初始化路由
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.",
				ballotX, err)
		}
		t.Log("\tShould be able to create a request.",
			checkMark)

		rw := httptest.NewRecorder()
		// 调用默认的多路选择器的 ServerHttp方法，模仿了外部客户端对 /sendjson 服务端点的请求
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an Email.", ballotX, u.Email)
		}
	}
}
