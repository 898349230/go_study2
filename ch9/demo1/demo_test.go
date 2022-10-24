package main

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// 基础单元测试, 文件名称需要以 _test 结尾， 测试 使用 go test -v 命令运行测试
// 函数名称要以 Test 开头，并且 函数的签名必须接收一个指向 testing.T 类型的指针，并且不返回任何值
func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200
	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				// Fatal() 会立刻停止当前这个测试函数的执行，但如果还有其他没有执行的测试函数会继续执行其他测试函数
				//  Error() 函数不会停止
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
