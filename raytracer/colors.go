package raytracer

import (
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

func (c *Color) Add(b *Color) *Color {
	return &Color{c.R + b.R, c.G + b.G, c.B + b.B}
}

func (c *Color) Sub(b *Color) *Color {
	return &Color{c.R - b.R, c.G - b.G, c.B - b.B}
}

func (c *Color) Mul(scalar float64) *Color {
	return &Color{c.R * scalar, c.G * scalar, c.B * scalar}
}

func (c *Color) Equals(b *Color) bool {
	if math.Abs(c.R-b.R) > epsilon {
		return false
	}

	if math.Abs(c.G-b.G) > epsilon {
		return false
	}

	if math.Abs(c.B-b.B) > epsilon {
		return false
	}

	return true
}

func (c *Color) Prod(b *Color) *Color {
	return &Color{c.R * b.R, c.G * b.G, c.B * b.B}
}

func (c *Color) RGBA() *color.RGBA {
	return &color.RGBA{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: 255,
	}
}
