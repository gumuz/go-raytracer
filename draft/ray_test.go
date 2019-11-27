package main

import (
	"math"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	tuple := Tuple{4.3, -4.2, 3.1, 1.0}

	if !tuple.IsPoint() {
		t.Errorf("%v is not a Point", tuple)
	}
}

func TestTupleIsVector(t *testing.T) {
	tuple := Tuple{4.3, -4.2, 3.1, 0}

	if !tuple.IsVector() {
		t.Errorf("%v is not a Vector", tuple)
	}
}

func TestNewPoint(t *testing.T) {
	point := NewPoint(4.3, -4.2, 3.1)
	tuple := &Tuple{4.3, -4.2, 3.1, 1}

	if !point.Equals(tuple) {
		t.Errorf("%v is incorrect", point)
	}
}

func TestNewVector(t *testing.T) {
	vector := NewVector(4.3, -4.2, 3.1)
	tuple := &Tuple{4.3, -4.2, 3.1, 0}

	if !vector.Equals(tuple) {
		t.Errorf("%v is incorrect", vector)
	}
}

func TestAdd(t *testing.T) {
	point := NewPoint(3, -2, 5)
	vector := NewVector(-2, 3, 1)

	newPoint := NewPoint(1, 1, 6)

	if !point.Add(vector).Equals(newPoint) {
		t.Errorf("%v is incorrect", point.Add(vector))
	}
}

func TestSubPoints(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	vector := NewVector(-2, -4, -6)

	if !point1.Sub(point2).Equals(vector) {
		t.Errorf("%v is incorrect", point1.Sub(point2))
	}
}

func TestSubVector(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)

	point2 := NewPoint(-2, -4, -6)

	if !point1.Sub(vector).Equals(point2) {
		t.Errorf("%v is incorrect", point1.Sub(vector))
	}
}

func TestSubVectors(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)

	vector := NewVector(-2, -4, -6)

	if !vector1.Sub(vector2).Equals(vector) {
		t.Errorf("%v is incorrect", vector1.Sub(vector2))
	}
}

func TestNegate(t *testing.T) {
	tuple := &Tuple{1, -2, 3, -4}
	negTuple := &Tuple{-1, 2, -3, 4}

	if !tuple.Neg().Equals(negTuple) {
		t.Errorf("%v is incorrect", tuple.Neg())
	}
}

func TestMult(t *testing.T) {
	tuple := &Tuple{1, -2, 3, -4}
	multTuple := &Tuple{0.5, -1, 1.5, -2}

	if !tuple.Mult(.5).Equals(multTuple) {
		t.Errorf("%v is incorrect", tuple.Mult(.5))
	}
}

func TestDiv(t *testing.T) {
	tuple := &Tuple{1, -2, 3, -4}
	divTuple := &Tuple{0.5, -1, 1.5, -2}

	if !tuple.Div(2).Equals(divTuple) {
		t.Errorf("%v is incorrect", tuple.Div(2))
	}
}

func TestMag(t *testing.T) {
	vector1 := NewVector(1, 0, 0)

	if vector1.Mag() != 1 {
		t.Errorf("%v is incorrect", vector1.Mag())
	}

	vector2 := NewVector(0, 1, 0)

	if vector2.Mag() != 1 {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

	vector3 := NewVector(0, 0, 1)

	if vector3.Mag() != 1 {
		t.Errorf("%v is incorrect", vector3.Mag())
	}

	vector4 := NewVector(1, 2, 3)

	if vector4.Mag() != math.Sqrt(14) {
		t.Errorf("%v is incorrect", vector4.Mag())
	}

	vector5 := NewVector(-1, -2, -3)

	if vector5.Mag() != math.Sqrt(14) {
		t.Errorf("%v is incorrect", vector5.Mag())
	}
}

func TestNorm(t *testing.T) {
	vector1 := NewVector(4, 0, 0)

	if !vector1.Norm().Equals(NewVector(1, 0, 0)) {
		t.Errorf("%v is incorrect", vector1.Mag())
	}

	vector2 := NewVector(1, 2, 3)

	if !vector2.Norm().Equals(NewVector(0.26726, 0.53452, 0.80178)) {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

	if vector2.Norm().Mag() != 1 {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

}

func TestDot(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	if vector1.Dot(vector2) != 20 {
		t.Errorf("%v is incorrect", vector1.Dot(vector2))
	}
}

func TestCross(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	if !vector1.Cross(vector2).Equals(NewVector(-1, 2, -1)) {
		t.Errorf("%v is incorrect", vector1.Cross(vector2))
	}

	if !vector2.Cross(vector1).Equals(NewVector(1, -2, 1)) {
		t.Errorf("%v is incorrect", vector1.Cross(vector2))
	}
}

func TestAddColors(t *testing.T) {
	color1 := NewColor(0.9, 0.6, 0.75)
	color2 := NewColor(0.7, 0.1, 0.25)

	newColor := NewColor(1.6, 0.7, 1.0)

	if !color1.Add(color2).Equals(newColor) {
		t.Errorf("%v is incorrect", color1.Add(color2))
	}
}

func TestSubColors(t *testing.T) {
	color1 := NewColor(0.9, 0.6, 0.75)
	color2 := NewColor(0.7, 0.1, 0.25)

	newColor := NewColor(0.2, 0.5, 0.5)

	if !color1.Sub(color2).Equals(newColor) {
		t.Errorf("%v is incorrect", color1.Sub(color2))
	}
}

func TestMultColor(t *testing.T) {
	color1 := NewColor(0.2, 0.3, 0.4)
	color2 := NewColor(0.4, 0.6, 0.8)

	if !color1.Mult(2).Equals(color2) {
		t.Errorf("%v is incorrect", color1.Mult(2))
	}
}

func TestProdColor(t *testing.T) {
	color1 := NewColor(1, 0.2, 0.4)
	color2 := NewColor(0.9, 1, 0.1)

	if !color1.Prod(color2).Equals(NewColor(0.9, 0.2, 0.04)) {
		t.Errorf("%v is incorrect", color1.Prod(color2))
	}
}
