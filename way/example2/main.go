// 本示例用来展示指针接受者的方法
package main

import "fmt"

// Point 表示点
type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pPtr := &p
	pPtr.ScaleBy(2)
	fmt.Println(p)

	p2 := Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2)

	p3 := Point{1, 2}
	p3.ScaleBy(2)
	fmt.Println(p3)

	//Point{1,2}.ScaleBy(2) // 编译错误: 不能获得Point类型的字面量的地址
	(&Point{1,2}).ScaleBy(2)
}
