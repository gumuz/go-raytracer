package raytracer

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (p *Vector) Equals(b *Vector) bool {
	return math.Abs(p.X-b.X) < epsilon && math.Abs(p.Y-b.Y) < epsilon && math.Abs(p.Z-b.Z) < epsilon
}

func (v *Vector) Add(b *Vector) *Vector {
	return NewVector(v.X+b.X, v.Y+b.Y, v.Z+b.Z)
}

func (v *Vector) Sub(b *Vector) *Vector {
	return NewVector(v.X-b.X, v.Y-b.Y, v.Z-b.Z)
}

func (v *Vector) Mult(scalar float64) *Vector {
	return NewVector(v.X*scalar, v.Y*scalar, v.Z*scalar)
}

func (v *Vector) Div(scalar float64) *Vector {
	return NewVector(v.X/scalar, v.Y/scalar, v.Z/scalar)
}

func (v *Vector) Neg() *Vector {
	return NewVector(0, 0, 0).Sub(v)
}

func (v *Vector) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v *Vector) Norm() *Vector {
	magnitude := v.Mag()
	return v.Div(magnitude)
}

func (v *Vector) Dot(b *Vector) float64 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

func (v *Vector) Cross(b *Vector) *Vector {
	return NewVector(
		v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X)
}
