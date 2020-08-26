package main
//export PATH=$PATH:/usr/local/go/bin
//copy and paste exer10 directory into home/go/src
//go test exer10 -v

import (
	"fmt"
	"math"
	"image"
	"image/png"
	"os"
)
 
//put in fibonacci.go
func Fibonacci(n uint) uint {
	return Fib(n, 0)
}

func Fibby(n uint) uint {
	if(n == 0){return 0}
	if(n == 1){return 1}
	return (Fibby(n-1) + Fibby(n-2))
}

func FibbyPar(n uint, cutoff uint, channel chan uint) {
	channel <- Fib(n-1,cutoff)
	channel <- Fib(n-2,cutoff)
}

func Fib(n uint, cutoff uint) uint {
	var fibNum uint = 0
	if(n < cutoff || n < 2){
		fibNum = Fibby(n)
	} else{
		fibChan := make(chan uint)
		go FibbyPar(n, cutoff, fibChan)
		fibNum = <-fibChan + <-fibChan
	}
	return fibNum
}

//copy points.go into exercise, change package to exer10
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

//put in triangle.go
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

type Transformable interface {
	Rotate(float64)
	Scale(int)
}

func TurnDouble(t Transformable, angle float64) {
	t.Scale(2)
	t.Rotate(angle)
}

//put in image.go
type Col struct {
    r, g, b, alpha uint8
}

func (c Col) RGBA() (r, g, b, a uint32) {
	r = uint32(c.r)
	r |= r << 8
	g = uint32(c.g)
	g |= g << 8
	b = uint32(c.b)
	b |= b << 8
	a = uint32(c.alpha)
	a |= a << 8
	return
}

func GetRadius(px, py, centerx, centery int) float64 {
	x := math.Abs(float64(px - centerx))
	y := math.Abs(float64(py - centery))
	return math.Sqrt(x*x + y*y)
}

func DrawCircle(outerRadius, innerRadius int, outputFile string) {
	width := 200
	height := 200
	topLeft := image.Point{0,0}
	bottomRight := image.Point{width,height}
	img := image.NewRGBA(image.Rectangle{topLeft,bottomRight})
	black := Col{0,0,0,0xff}
	white := Col{255, 255, 255, 0xff}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			radius := int(GetRadius(x,y,width/2,height/2))
			if  radius < outerRadius && radius > innerRadius{
				img.Set(x,y,black)
			} else {
				img.Set(x,y,white)
			}
		}
	}
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	png.Encode(file, img)
	defer file.Close()
}

//main
func main() {
	//Fibonacci test
	fibTest := Fibonacci(20)
	fmt.Println(fibTest)
	//Triangle test
	p1 := NewPoint(1,1)
	p2 := NewPoint(1,0)
	p3 := NewPoint(0,0)
	triangle := NewTriangle(p1,p2,p3)
	fmt.Println("original triangle is: ", triangle.String())
	triangle.Rotate(math.Pi / 2)
	fmt.Println("Rotate triangle is: ", triangle.String())
	triangle.Scale(5)
	fmt.Println("scaled triangle is: ", triangle.String())
	//Transformable test
	pt := Point{3, 4}
	TurnDouble(&pt, 3*math.Pi/2)
	fmt.Println(pt)
	tri := Triangle{Point{1, 2}, Point{-3, 4}, Point{5, -6}}
	TurnDouble(&tri, math.Pi)
	fmt.Println(tri)
	//image test
	DrawCircle(40, 20, "out.png")
}	