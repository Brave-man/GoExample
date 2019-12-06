package main

import "fmt"

func main() {
	// 数组是内存中一段连续的值类型一致的空间
	// 1.声明数组: 长度为5,值为零值
	var array [5]int
	fmt.Println(len(array))
	// 2.字面量声明数组
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(array2))
	// 3.自动计算数组的长度
	array3 := [...]int{88, 99, 77}
	fmt.Println(len(array3))
	// 4.声明数组并指定指定元素的值
	array4 := [5]int{1: 10, 3: 20}
	fmt.Println(array4)
}
