package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string // 类型再定义

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code: %s\n", traceCode)
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
	// WithValue返回父节点的副本，其中与key关联的值为val。
	// 仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
	// 所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
	// WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。
	// 或者，导出的上下文关键变量的静态类型应该是指针或接口.
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12112121")
	wg.Add(1)
	go worker(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("over")
}
