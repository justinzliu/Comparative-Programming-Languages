package exer10

import (
	"math"
	"image"
	"image/png"
	"os"
)

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