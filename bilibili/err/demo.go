package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// defer 确保所有 deferred 的动作可以在函数返回前执行， 可用于释放资源
	defer f.Close()
	// 写入文件
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check erros, handle them graceffuly")
	return err

}

type SudokuError []error

// 实现 Error 接口
func (se SudokuError) Error() string {

	return "self error"
}

func main() {
	err := proverbs("")
	if err != nil {
		// err.(SudokuError)  断言
		if errs, ok := err.(SudokuError); ok {
			fmt.Println("断言 true ", errs)
		} else {
			fmt.Println("断言 false ", errs)
		}
		os.Exit(1)
	}
}
