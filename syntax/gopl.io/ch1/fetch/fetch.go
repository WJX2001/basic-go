package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func functionDemo1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}

// TODO: practice1: 调用io.Copy(dst,src)会从src中读取内容放入dst中
func functionDemo2() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1: %v\n", err)
			os.Exit(1)
		}
		// 使用io.Copy 方法读取内容
		println("我看下")
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch2: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// TODO: practice2: 修改fetch这个案例，如果输入的url参数没有http://前缀，为这个url加上该前缀，可能会用到strings.HasPrefix这个函数
func functionDemo3() {
	for _, url := range os.Args[1:] {

		// 请求接口前，判断url是否包含http://前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		// 请求接口
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// 获取内容
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// TODO: 修改fetch 打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码
func functionDemo4() {
	for _, url := range os.Args[1:] {

		// 请求接口前，判断url是否包含http://前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		// 请求接口
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// 打印HTTP状态码
		fmt.Printf("这是状态码: %s\n", resp.Status)

		// 获取内容
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func main() {
	// functionDemo1()
	functionDemo4()
}
