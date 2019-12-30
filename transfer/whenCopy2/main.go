// 什么时候会发生副本创建
// 2.slice，map和数组在初始化和按索引设置的时候也会创建副本
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

	// slice
	s := []Bird{parrot1}
	s = append(s, parrot1)
	parrot1.Age = 3
	fmt.Printf("parrot2: %+v, ptr: %p\n", s[0], &(s[0]))
	fmt.Printf("parrot3: %+v, ptr: %p\n", s[1], &(s[1]))
	parrot1.Age = 1

	// map
	m := make(map[int]Bird)
	m[0] = parrot1
	parrot1.Age = 4
	fmt.Printf("parrot4: %+v\n", m[0])
	parrot1.Age = 5
	parrot5 := m[0]
	fmt.Printf("parrot5: %+v, ptr: %p\n", parrot5, &parrot5)
	parrot1.Age = 1

	// array
	a := [2]Bird{parrot1}
	parrot1.Age = 6
	fmt.Printf("parrot6: %+v, ptr: %p\n", a[0], &(a[0]))
	parrot1.Age = 1
	a[1] = parrot1
	parrot1.Age = 7
	fmt.Printf("parrot7: %+v, ptr: %p\n", a[1], &(a[1]))
}

//parrot1: {Age:1 Name:Blue}, ptr: 0x1174bf0
//parrot2: {Age:1 Name:Blue}, ptr: 0xc000076180
//parrot3: {Age:1 Name:Blue}, ptr: 0xc000076198
//parrot4: {Age:1 Name:Blue}
//parrot5: {Age:1 Name:Blue}, ptr: 0xc00008c0c0
//parrot6: {Age:1 Name:Blue}, ptr: 0xc0000761b0
//parrot7: {Age:1 Name:Blue}, ptr: 0xc0000761c8
// 可以看出: slice/map/数组 的元素全是原始变量的副本
