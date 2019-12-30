// 什么时候会发生副本创建
// 4.channel
// 往channel中send对象的时候也会创建对象的副本
package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

var parrot1 = Bird{
	Age:  1,
	Name: "Blue",
}

func main() {
	ch := make(chan Bird, 3)
	fmt.Printf("parrot1: %+v, ptr: %p\n", parrot1, &parrot1)
	ch <- parrot1
	parrot1.Age = 2
	ch <- parrot1
	parrot1.Age = 3
	ch <- parrot1
	close(ch)

	p := <-ch
	fmt.Printf("parrot%d: %+v, ptr: %p\n", 2, p, &p)
	p = <-ch
	fmt.Printf("parrot%d: %+v, ptr: %p\n", 3, p, &p)
	p = <-ch
	fmt.Printf("parrot%d: %+v, ptr: %p\n", 4, p, &p)
}

//parrot1: {Age:1 Name:Blue}, ptr: 0x1173bf0
//parrot2: {Age:1 Name:Blue}, ptr: 0xc00008c040
//parrot3: {Age:2 Name:Blue}, ptr: 0xc00008c040
//parrot4: {Age:3 Name:Blue}, ptr: 0xc00008c040
