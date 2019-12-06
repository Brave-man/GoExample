package main

import (
	"fmt"
	"runtime"
)

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在的函数
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动触发宕机前")
		panic(&panicContext{function:"手动触发panic"})
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
	})

	fmt.Println("运行后")
}
