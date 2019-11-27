package raytracer

import (
	"image"
	imgcolor "image/color"
	"image/png"
	"log"
	"os"
)

type canvas struct {
	width, height int
	img           *image.NRGBA
}

func NewCanvas(width, height int) *canvas {
	return &canvas{width, height, image.NewNRGBA(image.Rect(0, 0, width, height))}
}

func (c *canvas) WritePixel(x, y int, color *color) {
	c.img.Set(x, c.height-y, imgcolor.NRGBA{
		R: uint8(color.R * 255),
		G: uint8(color.G * 255),
		B: uint8(color.B * 255),
		A: 255,
	})
}

func (c *canvas) PixelAt(x, y int) *color {
	r, g, b, _ := c.img.At(x, y).RGBA()
	return NewColor(float64(r)/255, float64(g)/255, float64(b)/255)
}

func (c *canvas) Save() {
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
