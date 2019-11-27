package raytracer

import "math"

type Point struct {
	X, Y, Z float64
}

func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

func (p *Point) Equals(b *Point) bool {
	return math.Abs(p.X-b.X) < epsilon && math.Abs(p.Y-b.Y) < epsilon && math.Abs(p.Z-b.Z) < epsilon
}

func (p *Point) AddV(v *Vector) *Point {
	return NewPoint(p.X+v.X, p.Y+v.Y, p.Z+v.Z)
}

func (p *Point) Sub(b *Point) *Vector {
	return NewVector(p.X-b.X, p.Y-b.Y, p.Z-b.Z)
}

func (p *Point) SubV(v *Vector) *Point {
	return NewPoint(p.X-v.X, p.Y-v.Y, p.Z-v.Z)
}
