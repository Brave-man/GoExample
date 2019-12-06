package main

import "fmt"

func main() {
	// 声明nil切片
	var slice []int
	fmt.Println(len(slice), cap(slice))
	// 声明空切片
	slice1 := make([]int, 0)
	slice2 := []int{}
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	slice3 := []int{10, 20, 30, 40, 50}
	fmt.Println(len(slice3), cap(slice3))
	newSlice := slice3[1:3]
	fmt.Println(len(newSlice), cap(newSlice))
	newSlice[1] = 300
	fmt.Println(slice3[2])
}
