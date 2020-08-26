package main

import (
	"fmt"
	"math"
)

//put in hailstone.go
func Hailstone(x uint) uint {
	//even
	if (x%2 == 0) {
		return x/2
	}
	return ((3*x)+1)
}

func HailstoneSequenceAppend(x uint) []uint {
	start := []uint{}
	curr := x
	for curr > 1 {
		start = append(start,curr)
		curr = Hailstone(curr)
	}
	start = append(start,1)
	return start
}

func HailstoneLen(x uint) uint {
	curr := x
	var len uint = 1
	for curr > 1 {
		curr = Hailstone(curr)
		len++
	}
	return len
}

func HailstoneSequenceAllocate(x uint) []uint {
	len := HailstoneLen(x)
	start := make([]uint, len, len)
	curr := x
	it := 0
	for curr > 1 {
		start[it] = curr
		curr = Hailstone(curr)
		it++
	}
	start[len-1] = 1
	return start
}

//put in points.go
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

func main() {
	//Hailstone test
	fmt.Println(Hailstone(5))
	//HailstoneLen test
	fmt.Println(HailstoneLen(5))
	//HailstoneSequenceAppend test
	fmt.Println(HailstoneSequenceAppend(5))
	//HailstoneAllocate test
	fmt.Println(HailstoneSequenceAllocate(5))
	/*
	//String Representation test
	p := NewPoint(3,4.5)
	fmt.Println(p.String() == "(3, 4.5)")
	//Calculate Norm test
	pt := NewPoint(3, 4)
	fmt.Println(pt.Norm() == 5.0)
	*/
}	