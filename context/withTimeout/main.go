package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(10 * time.Millisecond) // 假设正常连接数据库耗时10ms
		select {
		case <-ctx.Done(): // 50ms后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	wg.Add(1)
	go worker(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("over")
}
