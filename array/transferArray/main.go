package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个8MB的数组
	var array [1000000]int
	// 将数组传递给函数
	now := time.Now()
	foo(array)
	fmt.Println("foo():", time.Since(now))

	now2 := time.Now()
	foo2(&array)
	fmt.Println("foo2()", time.Since(now2))
}

func foo(array [1000000]int) {
	fmt.Println(len(array))
}

func foo2(array *[1000000]int) {
	fmt.Println(len(*array))
}