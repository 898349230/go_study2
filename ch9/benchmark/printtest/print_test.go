// 基准测试的文件名也必须以 _test.go 结尾
package print_test

import (
	"fmt"
	"testing"
)

// 基准测试， 基准测试函数必须以 Benchmark 开头，接收一个指向 testing.B 类型的指针作为唯一参数
// 基准测试命令： go test -v -run="none" -bench="BenchmarkSprintf"
// 其他参数： -benchmem， 提供每次操作分配内存的次数，以及总共分配内存的字节数
// 			  -benchtime，更改测试执行的最短时间
func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer()

	// 基准测试框架默认会在持续1s的时间内反复调用需要测试的函数
	// 每次调用测试函数时，都会增加 b.N的值，第一次调用时，b.N 的值为1，
	// 需要注意的是一定要将所有要进行基准测试的代码都放到循环里，并且循环要使用 b.N 的值
	for i := 0; i < b.N; i++ {
		fmt.Printf("%d", number)
		// strconv.FormatInt(int64(number), 1)
		// strconv.Itoa(number)
	}
}
