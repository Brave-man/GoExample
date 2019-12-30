package main

import "fmt"

type Bird struct {
	Age  int
	Name string
}

func passV(b *Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:%+v, 内存地址:%p, 指针的内存地址:%p\n", *b, b, &b)
}

func main() {
	parrot := &Bird{
		Age:  1,
		Name: "Blue",
	}
	fmt.Printf("原始的Bird:%+v, 内存地址:%p, 指针的内存地址:%p\n", *parrot, parrot, &parrot)
	// 函数passV接受一个指针, 传入函数后，函数会创建指针的副本
	passV(parrot)
	fmt.Printf("调用后的Bird:%+v, 内存地址:%p, 指针的内存地址:%p\n", *parrot, parrot, &parrot)
}

//原始的Bird:{Age:1 Name:Blue}, 内存地址:0xc00008c020, 指针的内存地址:0xc000094018
//传入修改后的Bird:{Age:2 Name:GreatBlue}, 内存地址:0xc00008c020, 指针的内存地址:0xc000094028
//调用后的Bird:{Age:2 Name:GreatBlue}, 内存地址:0xc00008c020, 指针的内存地址:0xc000094018
