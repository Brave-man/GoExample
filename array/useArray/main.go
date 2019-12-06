package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	// 1.修改数组的元素
	array[2] = 35
	fmt.Println(array[2])

	// 2.访问指针数组的元素
	array2 := [5]*int{0: new(int), 1: new(int)}
	*array2[0] = 10
	*array2[1] = 20
	var num = 3
	array2[2] = &num
	fmt.Println(array2)

	// 3.把同样类型的数组赋值给另外一个数组
	var array3 [5]int
	array4 := [5]int{10, 20, 30, 40, 50}
	array3 = array4
	fmt.Println(array3)
	fmt.Println(array4)
	array4[1] = 22
	fmt.Println(array3)
	fmt.Println(array4)

	// 4.把一个指针数组赋值给另一个: 只会赋值指针的值，而不会复制指针所指向的值
	var array5 [3]*string
	array6 := [3]*string{new(string), new(string), new(string)}
	*array6[0] = "Red"
	*array6[1] = "Blue"
	*array6[2] = "Green"
	array5 = array6
	fmt.Println(array5)
	fmt.Println(array6)
}
