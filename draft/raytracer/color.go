package raytracer

import "math"

type color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) *color {
	return &color{r, g, b}
}

func (c *color) Add(b *color) *color {
	return NewColor(c.R+b.R, c.G+b.G, c.B+b.B)
}

func (c *color) Sub(b *color) *color {
	return NewColor(c.R-b.R, c.G-b.G, c.B-b.B)
}

func (c *color) Mult(scalar float64) *color {
	return NewColor(c.R*scalar, c.G*scalar, c.B*scalar)
}

func (c *color) Equals(b *color) bool {
	return math.Abs(c.R-b.R) < epsilon && math.Abs(c.G-b.G) < epsilon && math.Abs(c.B-b.B) < epsilon
}

func (c *color) Prod(b *color) *color {
	return NewColor(c.R*b.R, c.G*b.G, c.B*b.B)
}
