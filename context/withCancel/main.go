package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // 直接结束该goroutine，防止泄漏
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	// withCancel返回带有新的Done通道的父节点的副本，当调用返回的cancel函数或关闭父上下文的Done通道时，将关闭返回上下文的Done通道
	// 无论先发生什么情况。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当取完所需的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
