package raytracer

import (
	"image"
	"image/png"
	"log"
	"os"
)

type Canvas struct {
	image *image.RGBA
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{image.NewRGBA(image.Rect(0, 0, width, height))}
}

func (c *Canvas) SetAt(x, y int, color *Color) {
	c.image.Set(x, y, color.RGBA())
}

func (c *Canvas) GetAt(x, y int) *Color {
	r, g, b, _ := c.image.At(x, y).RGBA()
	return &Color{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

func (c *Canvas) Save(filename string) {
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, c.image); err != nil {
		log.Fatal(err)
	}

}
