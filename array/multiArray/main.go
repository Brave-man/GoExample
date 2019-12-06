package main

import (
	"fmt"
)

func main() {
	// 1.声明二维数组
	var array [4][2]int
	fmt.Println(array)

	// 2.使用数组字面量声明数组
	array1 := [4][2]int{{11, 22}, {111, 222}, {22, 222}, {33, 333}}
	fmt.Println(array1)

	// 3.声明并初始化数组部分元素
	array2 := [4][2]int{1: {22, 222}, 3:{333, 444}}
	fmt.Println(array2)

	// 4.声明并初始化二维数组的单个元素
	array3 := [4][2]int{1: {0: 222}, 3: {1: 444}}
	fmt.Println(array3)

	// 5.访问二维数组的元素
	var array5 [2][2]int
	array5[0][0] = 1
	array5[0][1] = 2
	array5[1][0] = 10
	array5[1][1] = 20

	// 6.复制二维数组到同类型二维数组中
	var array6 [2][2]int
	array6 = array5
	fmt.Println(array5)
	fmt.Println(array6)

	// 7.使用索引对多维数组赋值
	var array7 [2]int = array6[1]
	fmt.Println(array7)

	var value = array5[0][1]
	fmt.Println(value)
}
