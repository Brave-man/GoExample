package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

func passV(b Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p\n", b, &b)
}

func main() {
	parrot := Bird{
		Age:  1,
		Name: "Blue",
	}
	fmt.Printf("原始的Bird:\t %+v, \t内存地址：%p\n", parrot, &parrot)
	// 在T类型作为参数的时候，传递参数parrot，会将它的副本传递给passV, 在函数内对参数的修改会影响原来的对象
	passV(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t内存地址：%p\n", parrot, &parrot)
}

//原始的Bird:      {Age:1 Name:Blue},     内存地址：0xc00008c020
//传入修改后的Bird:        {Age:2 Name:GreatBlue},        内存地址：0xc00008c060
//调用后原始的Bird:        {Age:1 Name:Blue},     内存地址：0xc00008c020
