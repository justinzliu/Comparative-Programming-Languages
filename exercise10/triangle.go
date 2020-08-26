package exer10

import (
	"math"
)

//Point Structure
type Point struct {
	x  float64
	y  float64
}
func NewPoint(newx float64, newy float64) Point {
	var point Point
	point.x = newx
	point.y = newy
	return point
}
func (p Point) String() string {
	stringed := fmt.Sprintf("(%v, %v)", p.x, p.y)
	return stringed
}
func (p Point) Norm() float64 {
	return (math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2)))
}
func (p *Point) Scale(factor int) {
	p.x = p.x * float64(factor)
	p.y = p.y * float64(factor)
}
func (p *Point) Rotate(angle float64) {
	var x = p.x
	var y = p.y
	p.x = (x * math.Cos(angle)) - (y * math.Sin(angle)) 
	p.y = (x * math.Sin(angle)) + (y * math.Cos(angle))
}

//Triangle Structure
type Triangle struct {
	A, B, C Point
}
func NewTriangle(p1 Point, p2 Point, p3 Point) Triangle {
	var triangle Triangle
	triangle.A = p1
	triangle.B = p2
	triangle.C = p3
	return triangle
}
func (t Triangle) String() string {
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}
func (t* Triangle) Scale(factor int) {
	t.A.Scale(factor)
	t.B.Scale(factor)
	t.C.Scale(factor)
}
func (t* Triangle) Rotate(angle float64) {
	t.A.Rotate(angle)
	t.B.Rotate(angle)
	t.C.Rotate(angle)
}

//Transformable
type Transformable interface {
	Rotate(float64)
	Scale(int)
}
func TurnDouble(t Transformable, angle float64) {
	t.Scale(2)
	t.Rotate(angle)
}