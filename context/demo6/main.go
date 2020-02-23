package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx) // 在子goroutine中再开启一个goroutine，关闭新开启的子goroutine, 只需传入ctx
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级的通知
			fmt.Println("worker:", ctx.Err())
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级的通知
			fmt.Println("worker2:", ctx.Err())
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel() // 通知子goroutine结束

	wg.Wait()
	fmt.Println("over")
}
