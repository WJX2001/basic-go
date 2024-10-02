package main

import "time"

import (
	"fmt"
	"sync"
)

var total int
var wg sync.WaitGroup

// 加入读写锁
var lock sync.RWMutex

// 读
func read() {
	defer wg.Done()
	lock.RLock() // 如果只是读数据，那么这个锁不产生影响，但是如果读写同时发生的时候，那么有影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}

func main() {
	wg.Add(2)
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
}
