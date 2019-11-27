package raytracer_test

import (
	"testing"

	"github.com/gumuz/go-raytracer/raytracer"
)

func TestNewColor(t *testing.T) {
	color := raytracer.NewColor(4.3, -4.2, 3.1)

	if color.R != 4.3 || color.G != -4.2 || color.B != 3.1 {
		t.Errorf("%v is incorrect", color)
	}
}

func TestAddColors(t *testing.T) {
	color1 := raytracer.NewColor(0.9, 0.6, 0.75)
	color2 := raytracer.NewColor(0.7, 0.1, 0.25)

	newColor := raytracer.NewColor(1.6, 0.7, 1.0)

	if !color1.Add(color2).Equals(newColor) {
		t.Errorf("%v is incorrect", color1.Add(color2))
	}
}

func TestSubColors(t *testing.T) {
	color1 := raytracer.NewColor(0.9, 0.6, 0.75)
	color2 := raytracer.NewColor(0.7, 0.1, 0.25)

	newColor := raytracer.NewColor(0.2, 0.5, 0.5)

	if !color1.Sub(color2).Equals(newColor) {
		t.Errorf("%v is incorrect", color1.Sub(color2))
	}
}

func TestMultColor(t *testing.T) {
	color1 := raytracer.NewColor(0.2, 0.3, 0.4)
	color2 := raytracer.NewColor(0.4, 0.6, 0.8)

	if !color1.Mult(2).Equals(color2) {
		t.Errorf("%v is incorrect", color1.Mult(2))
	}
}

func TestProdColor(t *testing.T) {
	color1 := raytracer.NewColor(1, 0.2, 0.4)
	color2 := raytracer.NewColor(0.9, 1, 0.1)

	if !color1.Prod(color2).Equals(raytracer.NewColor(0.9, 0.2, 0.04)) {
		t.Errorf("%v is incorrect", color1.Prod(color2))
	}
}
