// 什么时候会发生副本创建
// 3.for-range 循环
// for-range循环也是将元素的副本赋值给循环变量，所以变量得到的是集合元素的副本
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
	fmt.Printf("parrot1: %+v, ptr: %p\n", parrot1, &parrot1)
	fmt.Println()
	// slice
	s := []Bird{parrot1, parrot1, parrot1}
	s[0].Age = 1
	s[1].Age = 2
	s[2].Age = 3
	parrot1.Age = 4

	for i, p := range s {
		fmt.Printf("parrot%d: %+v, ptr: %p\n", i+2, p, &p)
	}
	parrot1.Age = 1
	fmt.Println()

	// map
	m := make(map[int]Bird)
	parrot1.Age = 1
	m[0] = parrot1
	parrot1.Age = 2
	m[1] = parrot1
	parrot1.Age = 3
	m[2] = parrot1
	parrot1.Age = 4
	for k, v := range m {
		fmt.Printf("parrot%d: %+v, ptr: %p\n", k+2, v, &v)
	}
	parrot1.Age = 1
	fmt.Println()

	// array
	a := [...]Bird{parrot1, parrot1, parrot1}
	a[0].Age = 1
	a[1].Age = 2
	a[2].Age = 3
	parrot1.Age = 4
	for i, p := range a {
		fmt.Printf("parrot%d: %+v, ptr: %p\n", i+2, p, &p)
	}
}
