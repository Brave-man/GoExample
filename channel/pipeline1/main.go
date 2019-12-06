package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter 生产数据
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer 计算平方
	go func() {
		for {
			// 当一个通道关闭后，再向通道执行发送操作，回导致奔溃
			// 而，通道的接收方在读尽通道内的值后，从关闭的通道中读到值为零值
			// 通过ok的结果，可以判断是不是从关闭的通道中读的值

			// 或者使用for x := range naturals{ ... } 接受完最后一个值后关闭循环
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer 在主goroutine中执行打印操作
	for x := range squares {
		fmt.Println(x)
		time.Sleep(100 * time.Millisecond)
	}

	// 1.关于通道的关闭，不是必须的
	// 2.只有在需要通知接收方goroutine所有数据已经发送完毕时，才需要关闭通道
	// 3.通道也是可以通过垃圾回收器根据它是否可以访问它来决定是否回收它，而不是根据它是否关闭
	// 4.与关闭文件的close不同，文件的close是必须的
	// 5.试图关闭一个空通道或者已经关闭的通道会导致宕机
}
