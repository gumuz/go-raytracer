package raytracer_test

import (
	"testing"

	"github.com/gumuz/go-raytracer/raytracer"
)

func TestNewPoint(t *testing.T) {
	point := raytracer.NewPoint(4.3, -4.2, 3.1)

	if point.X != 4.3 || point.Y != -4.2 || point.Z != 3.1 {
		t.Errorf("%v is incorrect", point)
	}
}

func TestPointEquals(t *testing.T) {
	point1 := raytracer.NewPoint(4.3, -4.2, 3.1)
	point2 := raytracer.NewPoint(4.3, -4.2, 3.1)

	if !point1.Equals(point2) {
		t.Errorf("%v != %v", point1, point2)
	}
}

func TestPointAddV(t *testing.T) {
	point := raytracer.NewPoint(3, -2, 5)
	vector := raytracer.NewVector(-2, 3, 1)

	result := point.AddV(vector)
	expected := raytracer.NewPoint(1, 1, 6)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestPointSubV(t *testing.T) {
	point := raytracer.NewPoint(3, 2, 1)
	vector := raytracer.NewVector(5, 6, 7)

	result := point.SubV(vector)
	expected := raytracer.NewPoint(-2, -4, -6)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestPointSub(t *testing.T) {
	point1 := raytracer.NewPoint(3, 2, 1)
	point2 := raytracer.NewPoint(5, 6, 7)

	result := point1.Sub(point2)
	expected := raytracer.NewVector(-2, -4, -6)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}
