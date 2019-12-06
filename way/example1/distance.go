// 本示例用来展示值接受者的方法
package main

import (
	"fmt"

	"goCode/review/way/point"
)


// Path 标示线段
type Path []point.Point

// Distance 计算线段的长度
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := point.Point{X: 1, Y: 2}
	q := point.Point{X: 2, Y: 4}
	fmt.Println(point.Distance(p, q))
	fmt.Println(p.Distance(q))

	path := Path{p, q}
	fmt.Println(path.Distance())
}
