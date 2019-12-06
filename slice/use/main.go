package main

import "fmt"

func main() {
	// ----使用切片向append中添加元素----
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(len(slice), cap(slice))
	newSlice := slice[1:3]
	fmt.Println(len(newSlice), cap(newSlice))
	fmt.Println()

	newSlice = append(newSlice, 60)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	fmt.Println(slice) // [10 20 30 60 50]
	fmt.Println()

	// ----使用append同时增加切片的长度和容量----
	slice2 := []int{10, 20, 30, 40}
	fmt.Println(len(slice2), cap(slice2))
	newSlice2 := append(slice2, 60)
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(newSlice2, len(newSlice2), cap(newSlice2))
	fmt.Println()

	// ----创建切片时的三个索引----
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice3 := source[2:3:4]
	fmt.Println(len(slice3), cap(slice3))
	//slice4 := source[2:3:6] // 超出容量时会抛出异常
	//fmt.Println(slice4)

	// ----设置长度和容量一致-----
	source2 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice4 := source2[2:3:3]
	slice4 = append(slice4, "Kiwi")
	fmt.Println(source2[3])
	fmt.Println(slice4[1])
	fmt.Println()

	source3 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice5 := source3[2:3]
	fmt.Println(len(slice5), cap(slice5))
	slice5 = append(slice5, "Kiwi")
	fmt.Println(source3[3])
	fmt.Println(slice5[1])

	// ----将一个切片追加到另外一个切片----
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println(append(s1, s2...))
}
