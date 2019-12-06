package main

import (
	"fmt"

	"goCode/review/way/point"
)

type Path []point.Point

func (path Path) TranslateBy(offset point.Point, add bool) {
	var op func(p, q point.Point) point.Point
	if add {
		op = point.Point.Add
	} else {
		op = point.Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := point.Point{X: 1, Y: 2}
	q := point.Point{X: 3, Y: 8}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	var origin point.Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	fmt.Println(p)

	distance := point.Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*point.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
