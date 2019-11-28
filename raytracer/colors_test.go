package raytracer_test

import (
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestColorTuple(t *testing.T) {
	/* Scenario: Colors are (red, green, blue) tuples
	   Given c ← color(-0.5, 0.4, 1.7)
	   Then c.red = -0.5
	     And c.green = 0.4
		 And c.blue = 1.7 */

	c := &rt.Color{-0.5, 0.4, 1.7}

	if c.R != -0.5 || c.G != 0.4 || c.B != 1.7 {
		t.Errorf("Error: %v", c)
	}

}

func TestAddColors(t *testing.T) {
	/* Scenario: Adding colors
	   Given c1 ← color(0.9, 0.6, 0.75)
	     And c2 ← color(0.7, 0.1, 0.25)
		Then c1 + c2 = color(1.6, 0.7, 1.0) */
	c1 := &rt.Color{0.9, 0.6, 0.75}
	c2 := &rt.Color{0.7, 0.1, 0.25}

	result := c1.Add(c2)

	expected := &rt.Color{1.6, 0.7, 1.0}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestSubColors(t *testing.T) {
	/* Scenario: Subtracting colors
	   Given c1 ← color(0.9, 0.6, 0.75)
	     And c2 ← color(0.7, 0.1, 0.25)
		Then c1 - c2 = color(0.2, 0.5, 0.5) */
	c1 := &rt.Color{0.9, 0.6, 0.75}
	c2 := &rt.Color{0.7, 0.1, 0.25}

	result := c1.Sub(c2)

	expected := &rt.Color{0.2, 0.5, 0.5}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMulColorsByScalar(t *testing.T) {
	/* Scenario: Multiplying a color by a scalar
	   Given c ← color(0.2, 0.3, 0.4)
	   Then c * 2 = color(0.4, 0.6, 0.8) */
	c := &rt.Color{0.2, 0.3, 0.4}

	result := c.Mul(2)

	expected := &rt.Color{0.4, 0.6, 0.8}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
func TestProdColors(t *testing.T) {
	/* Scenario: Multiplying colors
	   Given c1 ← color(1, 0.2, 0.4)
	     And c2 ← color(0.9, 1, 0.1)
		Then c1 * c2 = color(0.9, 0.2, 0.04) */
	c1 := &rt.Color{1, 0.2, 0.4}
	c2 := &rt.Color{0.9, 1, 0.1}

	result := c1.Prod(c2)

	expected := &rt.Color{0.9, 0.2, 0.04}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
