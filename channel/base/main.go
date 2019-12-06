package main

import "fmt"

func main() {
	// 1.channel是一个引用类型的数据，当复制或者作为参数传递的时候，复制的是引用
	// 因此调用者和接受者引用的是同一份数据结构
	// 通道的零值是nil
	ch := make(chan int) // 创建一个无缓冲通道
	ch2 := ch
	// 2.同种类型的通道可以做比较
	fmt.Println(ch == ch2)

	// 3.nil通道
	var ch3 chan int // nil
	fmt.Println(ch3)

	// 4.创建缓冲通道
	ch4 := make(chan int, 0) // 创建一个无缓冲通道
	ch5 := make(chan int, 10) // 创建一个容量是10的缓冲通道
	fmt.Println(ch4, ch5)

	// 5.通道的操作
	x := 5
	ch5 <- x // 向通道发送数据
	xx := <-ch5 // 从通道接受数据
	fmt.Println(x == xx)
	close(ch5) // 关闭通道
}
