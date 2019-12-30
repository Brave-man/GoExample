// 什么时候会发生副本创建
// 1.变量的赋值
// 包括函数外，函数内
package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

type Parrot struct {
	Age  int
	Name string
}

var parrot1 = Bird{
	Age:  1,
	Name: "Blue",
}
var parrot2 = parrot1

func main() {
	fmt.Printf("parrot1: %+v, ptr: %p\n", parrot1, &parrot1)
	fmt.Printf("parrot2: %+v, ptr: %p\n", parrot2, &parrot2)

	parrot3 := parrot1
	fmt.Printf("parrot3: %+v, ptr: %p\n", parrot3, &parrot3)

	parrot4 := Parrot(parrot1)
	fmt.Printf("parrot4: %+v, ptr: %p\n", parrot4, &parrot4)
}

//parrot1: {Age:1 Name:Blue}, ptr: 0x1173bf0
//parrot2: {Age:1 Name:Blue}, ptr: 0x1173c10
//parrot3: {Age:1 Name:Blue}, ptr: 0xc00000c0a0
//parrot4: {Age:1 Name:Blue}, ptr: 0xc00000c0e0
