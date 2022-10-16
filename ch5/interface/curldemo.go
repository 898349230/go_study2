package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Print("args error")
		os.Exit(-1)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Print(err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
	if err := resp.Body.Close(); err != nil {
		fmt.Print(err)
	}
}
