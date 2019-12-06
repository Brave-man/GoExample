package main

import (
	"fmt"
	"image/color"

	"goCode/review/way/point"
)

type ColoredPoint struct {
	point.Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.X)
	fmt.Println(cp.Point.X)

	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	var p = ColoredPoint{point.Point{X: 1, Y: 2}, red}
	var q = ColoredPoint{point.Point{X: 2, Y: 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	type ColoredPoint2 struct {
		*point.Point
		color.RGBA
	}

	p1 := ColoredPoint2{
		Point: &point.Point{X: 1, Y: 2},
		RGBA:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
	}
	q1 := ColoredPoint2{
		Point: &point.Point{X: 3, Y: 8},
		RGBA:  color.RGBA{R: 0, G: 255, B: 0, A: 255},
	}

	fmt.Println(p1.Distance(*q1.Point))
	q1.Point = p1.Point
	p1.ScaleBy(2)
	fmt.Println(*p1.Point, *q1.Point)
}
