package point

import "math"

// Point 表示点
type Point struct {
	X, Y float64
}

// Distance 普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 类型Point的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
