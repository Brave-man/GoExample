// 通过通道的方式关闭子goroutine
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
func worker(exit chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exit: // 通道关闭后，接受方不再阻塞，会接受到零值
			break LOOP // 退出LOOP循环
		default:
		}
	}
	wg.Done()
}

func main() {
	exit := make(chan struct{})

	wg.Add(1)
	go worker(exit)

	time.Sleep(3 * time.Second)
	close(exit) // 通过关闭通道，子goroutine就会接受到退出信号

	wg.Wait()
	fmt.Println("over")
}
