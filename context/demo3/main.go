// 通过全局变量的方式退出子goroutine
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

// 全局变量方式存在的问题:
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。
func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	// 如何接受外部命令实现退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的结束子goroutine
	time.Sleep(3*time.Second)
	exit = true // 通过修改全局变量实现子goroutine退出
	wg.Wait()
	fmt.Println("over")
}
