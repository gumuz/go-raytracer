package raytracer_test

import (
	"math"
	"testing"

	"github.com/gumuz/go-raytracer/raytracer"
)

func TestNewVector(t *testing.T) {
	vector := raytracer.NewVector(4.3, -4.2, 3.1)

	if vector.X != 4.3 || vector.Y != -4.2 || vector.Z != 3.1 {
		t.Errorf("%v is incorrect", vector)
	}
}

func TestVectorAdd(t *testing.T) {
	vector1 := raytracer.NewVector(3, -2, 5)
	vector2 := raytracer.NewVector(-2, 3, 1)

	result := vector1.Add(vector2)
	expected := raytracer.NewVector(1, 1, 6)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestVectorEquals(t *testing.T) {
	vector1 := raytracer.NewVector(4.3, -4.2, 3.1)
	vector2 := raytracer.NewVector(4.3, -4.2, 3.1)

	if !vector1.Equals(vector2) {
		t.Errorf("%v != %v", vector1, vector2)
	}
}

func TestVectorSub(t *testing.T) {
	vector1 := raytracer.NewVector(3, 2, 1)
	vector2 := raytracer.NewVector(5, 6, 7)

	result := vector1.Sub(vector2)
	expected := raytracer.NewVector(-2, -4, -6)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestVectorNegate(t *testing.T) {
	vector := raytracer.NewVector(1, -2, 3)

	result := vector.Neg()
	expected := raytracer.NewVector(-1, 2, -3)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestVectorMult(t *testing.T) {
	vector := raytracer.NewVector(1, -2, 3)

	result := vector.Mult(0.5)
	expected := raytracer.NewVector(0.5, -1, 1.5)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestVectorDiv(t *testing.T) {
	vector := raytracer.NewVector(1, -2, 3)

	result := vector.Div(2)
	expected := raytracer.NewVector(0.5, -1, 1.5)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestVectorMag(t *testing.T) {
	vector1 := raytracer.NewVector(1, 0, 0)

	if vector1.Mag() != 1 {
		t.Errorf("%v is incorrect", vector1.Mag())
	}

	vector2 := raytracer.NewVector(0, 1, 0)

	if vector2.Mag() != 1 {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

	vector3 := raytracer.NewVector(0, 0, 1)

	if vector3.Mag() != 1 {
		t.Errorf("%v is incorrect", vector3.Mag())
	}

	vector4 := raytracer.NewVector(1, 2, 3)

	if vector4.Mag() != math.Sqrt(14) {
		t.Errorf("%v is incorrect", vector4.Mag())
	}

	vector5 := raytracer.NewVector(-1, -2, -3)

	if vector5.Mag() != math.Sqrt(14) {
		t.Errorf("%v is incorrect", vector5.Mag())
	}
}

func TestVectorNorm(t *testing.T) {
	vector1 := raytracer.NewVector(4, 0, 0)

	if !vector1.Norm().Equals(raytracer.NewVector(1, 0, 0)) {
		t.Errorf("%v is incorrect", vector1.Mag())
	}

	vector2 := raytracer.NewVector(1, 2, 3)

	if !vector2.Norm().Equals(raytracer.NewVector(0.26726, 0.53452, 0.80178)) {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

	if vector2.Norm().Mag() != 1 {
		t.Errorf("%v is incorrect", vector2.Mag())
	}

}

func TestVectorDot(t *testing.T) {
	vector1 := raytracer.NewVector(1, 2, 3)
	vector2 := raytracer.NewVector(2, 3, 4)

	if vector1.Dot(vector2) != 20 {
		t.Errorf("%v is incorrect", vector1.Dot(vector2))
	}
}

func TestVectorCross(t *testing.T) {
	vector1 := raytracer.NewVector(1, 2, 3)
	vector2 := raytracer.NewVector(2, 3, 4)

	if !vector1.Cross(vector2).Equals(raytracer.NewVector(-1, 2, -1)) {
		t.Errorf("%v is incorrect", vector1.Cross(vector2))
	}

	if !vector2.Cross(vector1).Equals(raytracer.NewVector(1, -2, 1)) {
		t.Errorf("%v is incorrect", vector1.Cross(vector2))
	}
}
