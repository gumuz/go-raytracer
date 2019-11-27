package main

import (
	"image"
	imgcolor "image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const epsilon = 0.00001

type Color struct {
	r, g, b float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func (c *Color) Add(other *Color) *Color {
	return NewColor(c.r+other.r, c.g+other.g, c.b+other.b)
}

func (c *Color) Sub(other *Color) *Color {
	return NewColor(c.r-other.r, c.g-other.g, c.b-other.b)
}

func (c *Color) Mult(scalar float64) *Color {
	return NewColor(c.r*scalar, c.g*scalar, c.b*scalar)
}

func (c *Color) Equals(other *Color) bool {
	return math.Abs(c.r-other.r) < epsilon && math.Abs(c.g-other.g) < epsilon && math.Abs(c.b-other.b) < epsilon
}

func (c *Color) Prod(other *Color) *Color {
	return NewColor(c.r*other.r, c.g*other.g, c.b*other.b)
}

type Tuple struct {
	x, y, z, w float64
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0}
}

func (t *Tuple) IsPoint() bool {
	return t.w == 1
}

func (t *Tuple) IsVector() bool {
	return t.w == 0
}

func (t *Tuple) Equals(other *Tuple) bool {
	return math.Abs(t.x-other.x) < epsilon && math.Abs(t.y-other.y) < epsilon && math.Abs(t.z-other.z) < epsilon && math.Abs(t.w-other.w) < epsilon
}

func (t *Tuple) Add(other *Tuple) *Tuple {
	return &Tuple{t.x + other.x, t.y + other.y, t.z + other.z, t.w + other.w}
}

func (t *Tuple) Sub(other *Tuple) *Tuple {
	return &Tuple{t.x - other.x, t.y - other.y, t.z - other.z, t.w - other.w}
}

func (t *Tuple) Mult(scalar float64) *Tuple {
	return &Tuple{t.x * scalar, t.y * scalar, t.z * scalar, t.w * scalar}
}

func (t *Tuple) Div(scalar float64) *Tuple {
	return &Tuple{t.x / scalar, t.y / scalar, t.z / scalar, t.w / scalar}
}

func (t *Tuple) Neg() *Tuple {
	return (&Tuple{0, 0, 0, 0}).Sub(t)
}

func (t *Tuple) Mag() float64 {
	return math.Sqrt((t.x * t.x) + (t.y * t.y) + (t.z * t.z) + (t.w * t.w))
}

func (t *Tuple) Norm() *Tuple {
	return t.Div(t.Mag())
}

func (t *Tuple) Dot(other *Tuple) float64 {
	return t.x*other.x + t.y*other.y + t.z*other.z + t.w*other.w
}

func (t *Tuple) Cross(other *Tuple) *Tuple {
	return NewVector(
		t.y*other.z-t.z*other.y,
		t.z*other.x-t.x*other.z,
		t.x*other.y-t.y*other.x)
}

type Canvas struct {
	width, height int
	img           *image.NRGBA
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{width, height, image.NewNRGBA(image.Rect(0, 0, width, height))}
}

func (c *Canvas) WritePixel(x, y int, color *Color) {
	c.img.Set(x, y, imgcolor.NRGBA{
		R: uint8(color.r * 255),
		G: uint8(color.g * 255),
		B: uint8(color.b * 255),
		A: 255,
	})
}

func (c *Canvas) PixelAt(x, y int) *Color {
	r, g, b, _ := c.img.At(x, y).RGBA()
	return NewColor(float64(r)/255, float64(g)/255, float64(b)/255)
}

func (c *Canvas) Save() {
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, c.img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
