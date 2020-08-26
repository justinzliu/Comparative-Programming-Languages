package exer8

import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods
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
	//fmt.Println(stringed)
	return stringed
}

func (p Point) Norm() float64 {
	return (math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2)))
}