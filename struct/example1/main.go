package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 9
	w.Radius = 5
	w.Spokes = 20
	fmt.Printf("%+v\n", w)

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 9},
			Radius: 5,
		},
		Spokes: 20,
	}
	// 当匿名结构体中的所有成员是可比较的，那么结构体就是可比较的
	fmt.Println(w == w2)
}
