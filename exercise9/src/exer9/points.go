package exer9

// TODO: Point (with everything from exercise 8 and) and methods that modify them in-place

import (
	"fmt"
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

//Scale
func (p *Point) Scale(factor int) {
	p.x = p.x * float64(factor)
	p.y = p.y * float64(factor)
}

//Rotate
func (p *Point) Rotate(angle float64) {
	var x = p.x
	var y = p.y
	p.x = (x * math.Cos(angle)) - (y * math.Sin(angle)) 
	p.y = (x * math.Sin(angle)) + (y * math.Cos(angle))
}