package raytracer

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0}
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

func (t *Tuple) IsVector() bool {
	return t.W == 0
}

func (t *Tuple) Equals(b *Tuple) bool {
	if math.Abs(t.X-b.X) > epsilon {
		return false
	}

	if math.Abs(t.Y-b.Y) > epsilon {
		return false
	}

	if math.Abs(t.Z-b.Z) > epsilon {
		return false
	}

	if math.Abs(t.W-b.W) > epsilon {
		return false
	}

	return true
}

func (t *Tuple) Add(b *Tuple) *Tuple {
	return &Tuple{t.X + b.X, t.Y + b.Y, t.Z + b.Z, t.W + b.W}
}

func (t *Tuple) Sub(b *Tuple) *Tuple {
	return &Tuple{t.X - b.X, t.Y - b.Y, t.Z - b.Z, t.W - b.W}
}

func (t *Tuple) Neg() *Tuple {
	return &Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t *Tuple) Mul(scalar float64) *Tuple {
	return &Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar, t.W * scalar}
}

func (t *Tuple) Div(scalar float64) *Tuple {
	return &Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar, t.W / scalar}
}

func (t *Tuple) Mag() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

func (t *Tuple) Norm() *Tuple {
	return t.Div(t.Mag())
}

func (t *Tuple) Dot(b *Tuple) float64 {
	return t.X*b.X + t.Y*b.Y + t.Z*b.Z + t.W*b.W
}

func (t *Tuple) Cross(b *Tuple) *Tuple {
	return NewVector(
		t.Y*b.Z-t.Z*b.Y,
		t.Z*b.X-t.X*b.Z,
		t.X*b.Y-t.Y*b.X)

}

func (t *Tuple) Reflect(n *Tuple) *Tuple {
	return t.Sub(n.Mul(2).Mul(t.Dot(n)))

}
