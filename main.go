package main

import (
	"image/color"
	"image/png"
	"log"
	"os"
)

// Unnormalized color
type Color struct {
	R, B, G float64
}

type Image struct {
	cm     color.Model
	bounds image.Rectangle
	pixels [][]image.Color
}

func (i *Image) At(x, y int) color.Color {
	return i.pixels[x][y]
}

func (i *Image) Bounds() image.Rectangle {
	return i.bounds
}

func (i *Image) ColorModel() color.Model {
	return i.cm
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("No filename")
	}
	s := &Scene{Light: PointLight}
	file, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalln("Could not open file: %s", err.Error())
	}
	w, h := 200, 200
	image := Image{cm: color.RGBA}
	if err := png.Encode(file, image); err != nil {
		log.Fatalln("Error encoding: %s", err.Error())
	}
}

// x, y are from 0 to 1
func trace(s *Scene, v View, x, y float64) Color {
	r := getRay(v, x, y)
	for _, o := range s.Objects {
		p, d := o.Intersect(r)
	}
}

func getRay(v View, x, y float64) Ray {
	x = v.Rect.Width() * x
	y = v.Rect.Height() * y
	n := r.Rect.Normal()
	p := r.Rect.BottomCorner()
}

type Light interface {
	Intensity(p Point) float64
	To(p Point) Point
}

type Object interface {
	Intersect(r Ray) (Point, float64)
	Normal(p Point) Point
}

type Point struct {
	X, Y, Z float64
}

type Sphere struct {
	Center Point
	Radius float64
	Color  color.RGBA
}

type Ray struct {
	Origin, Dir Point
}

type View struct {
	Eye    Point
	Window Rect
}

type Rect interface {
	Plane
	BottomCorner() Point
	Length() float64
	Width() float64
}

type Plane interface {
	Normal() Point
}

type Scene struct {
	Light   Light
	Objects []Object
	Background
}

type PointLight struct {
	Origin Point
}

func (s *Sphere) Intersect(r Ray) (Point, float64) {

}

func (s *Sphere) Normal(p Point) Point {
	return p.Sub(s.Center)
}

func (l PointLight) Intensity(p Point) float64 {
	return p.Point.Dist(p)
}

func (l PointLight) To(p Point) Point {

}

func (a Point) Dist(b Point) float64 {
	return sqrt(sq(a.X-b.X) + sq(a.Y-b.Y) + sq(a.Z-a.Z))
}

func (a Point) Dot(b point) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Point) Sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Point) To(b Point) Point {
	return b.Sub(a)
}

func (a Point) Normalize() Point {
	total := a.Dot(a)
	return Point{a.X / total, a.Y / total, a.Z / total}
}

func (a Point) NormalizeWithTotal(total float64) Point {
	return Point{a.X / total, a.Y / total, a.Z / total}
}

func sq(x float64) float64 {
	return x * x
}

func computeShading(l Light, p Point, normal Point) float64 {
	baseLight := l.Intensity(p)
	return normal.Dot(l.To(p))
}

func (c Color) Color() color.RGBA {

}
