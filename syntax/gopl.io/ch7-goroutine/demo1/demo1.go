package demo1

import (
	"fmt"
	"strconv"
	"time"
)

/*
	1、在主线程中，开启一个goroutine，该goroutine每隔1秒输出 hello golang
	2、在主线程中也每隔一秒输出 hello msb，输出10次后，退出程序
	3、要求主线程和goroutine同时执行
*/

func Demo1() {
	for i := 0; i <= 100; i++ {
		fmt.Println("hello golang + " + strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(1 * time.Second)
	}
}
