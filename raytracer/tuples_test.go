package raytracer_test

import (
	"math"
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestPointTuple(t *testing.T) {
	/* Scenario: A tuple with w=1.0 is a point
	   Given a ← tuple(4.3, -4.2, 3.1, 1.0)
	   Then a.x = 4.3
	     And a.y = -4.2
	     And a.z = 3.1
	     And a.w = 1.0
	     And a is a point
		 And a is not a vector */
	a := &rt.Tuple{4.3, -4.2, 3.1, 1.0}

	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Errorf("Error: %v", a)
	}

	if !a.IsPoint() || a.IsVector() {
		t.Errorf("Error: %v %v", a.IsPoint(), a.IsVector())
	}
}

func TestVectorTuple(t *testing.T) {
	/* Scenario: A tuple with w=0 is a vector
	   Given a ← tuple(4.3, -4.2, 3.1, 0.0)
	   Then a.x = 4.3
	     And a.y = -4.2
	     And a.z = 3.1
	     And a.w = 0.0
	     And a is not a point
		 And a is a vector */
	a := &rt.Tuple{4.3, -4.2, 3.1, 0.0}

	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 0.0 {
		t.Errorf("Error: %v", a)
	}

	if a.IsPoint() || !a.IsVector() {
		t.Errorf("Error: %v %v", a.IsPoint(), a.IsVector())
	}
}

func TestCreatePointTuple(t *testing.T) {
	/* Scenario: point() creates tuples with w=1
	   Given p ← point(4, -4, 3)
	   Then p = tuple(4, -4, 3, 1) */
	p := rt.NewPoint(4, -4, 3)

	expected := &rt.Tuple{4, -4, 3, 1}
	if !p.Equals(expected) {
		t.Errorf("Error: %v", p)
	}
}

func TestCreateVectorTuple(t *testing.T) {
	/* Scenario: vector() creates tuples with w=0
	   Given v ← vector(4, -4, 3)
	   Then v = tuple(4, -4, 3, 0) */
	v := rt.NewVector(4, -4, 3)

	expected := &rt.Tuple{4, -4, 3, 0}
	if !v.Equals(expected) {
		t.Errorf("Error: %v", v)
	}
}

func TestAdd2Tuples(t *testing.T) {
	/* Scenario: Adding two tuples
	   Given a1 ← tuple(3, -2, 5, 1)
	     And a2 ← tuple(-2, 3, 1, 0)
		Then a1 + a2 = tuple(1, 1, 6, 1) */
	a1 := &rt.Tuple{3, -2, 5, 1}
	a2 := &rt.Tuple{-2, 3, 1, 0}

	result := a1.Add(a2)

	expected := &rt.Tuple{1, 1, 6, 1}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestSub2Points(t *testing.T) {
	/* Scenario: Subtracting two points
	   Given p1 ← point(3, 2, 1)
	     And p2 ← point(5, 6, 7)
	   Then p1 - p2 = vector(-2, -4, -6) */
	p1 := rt.NewPoint(3, 2, 1)
	p2 := rt.NewPoint(5, 6, 7)

	result := p1.Sub(p2)

	expected := rt.NewVector(-2, -4, -6)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestSubVectorFromPoint(t *testing.T) {
	/* Scenario: Subtracting a vector from a point
	   Given p ← point(3, 2, 1)
	     And v ← vector(5, 6, 7)
	   Then p - v = point(-2, -4, -6) */
	p := rt.NewPoint(3, 2, 1)
	v := rt.NewVector(5, 6, 7)

	result := p.Sub(v)

	expected := rt.NewPoint(-2, -4, -6)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestSub2Vectors(t *testing.T) {
	/* Scenario: Subtracting two vectors
	   Given v1 ← vector(3, 2, 1)
	     And v2 ← vector(5, 6, 7)
	   Then v1 - v2 = vector(-2, -4, -6) */
	v1 := rt.NewVector(3, 2, 1)
	v2 := rt.NewVector(5, 6, 7)

	result := v1.Sub(v2)

	expected := rt.NewVector(-2, -4, -6)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestSubVectorFromZeroVector(t *testing.T) {
	/* Scenario: Subtracting a vector from the zero vector
	   Given zero ← vector(0, 0, 0)
	     And v ← vector(1, -2, 3)
	   Then zero - v = vector(-1, 2, -3) */
	zero := rt.NewVector(0, 0, 0)
	v := rt.NewVector(1, -2, 3)

	result := zero.Sub(v)

	expected := rt.NewVector(-1, 2, -3)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestNegateTuple(t *testing.T) {
	/* Scenario: Negating a tuple
	   Given a ← tuple(1, -2, 3, -4)
	   Then -a = tuple(-1, 2, -3, 4) */
	a := &rt.Tuple{1, -2, 3, -4}

	result := a.Neg()

	expected := &rt.Tuple{-1, 2, -3, 4}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMulTupleByScalar(t *testing.T) {
	/* Scenario: Multiplying a tuple by a scalar
	   Given a ← tuple(1, -2, 3, -4)
	   Then a * 3.5 = tuple(3.5, -7, 10.5, -14) */
	a := &rt.Tuple{1, -2, 3, -4}

	result := a.Mul(3.5)

	expected := &rt.Tuple{3.5, -7, 10.5, -14}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMulTupleByFraction(t *testing.T) {
	/* Scenario: Multiplying a tuple by a fraction
	   Given a ← tuple(1, -2, 3, -4)
	   Then a * 0.5 = tuple(0.5, -1, 1.5, -2) */
	a := &rt.Tuple{1, -2, 3, -4}

	result := a.Mul(0.5)

	expected := &rt.Tuple{0.5, -1, 1.5, -2}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestDivTupleByScalar(t *testing.T) {
	/* Scenario: Dividing a tuple by a scalar
	   Given a ← tuple(1, -2, 3, -4)
	   Then a / 2 = tuple(0.5, -1, 1.5, -2) */
	a := &rt.Tuple{1, -2, 3, -4}

	result := a.Div(2)

	expected := &rt.Tuple{0.5, -1, 1.5, -2}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMagVectorX(t *testing.T) {
	/* Scenario: Computing the magnitude of vector(1, 0, 0)
	   Given v ← vector(1, 0, 0)
	   Then magnitude(v) = 1 */
	v := rt.NewVector(1, 0, 0)

	if v.Mag() != 1 {
		t.Errorf("Error: %v", v.Mag())
	}
}

func TestMagVectorY(t *testing.T) {
	/* Scenario: Computing the magnitude of vector(0, 1, 0)
	   Given v ← vector(0, 1, 0)
	   Then magnitude(v) = 1 */
	v := rt.NewVector(0, 1, 0)

	if v.Mag() != 1 {
		t.Errorf("Error: %v", v.Mag())
	}
}

func TestMagVectorZ(t *testing.T) {
	/* Scenario: Computing the magnitude of vector(0, 0, 1)
	   Given v ← vector(0, 0, 1)
	   Then magnitude(v) = 1 */
	v := rt.NewVector(0, 0, 1)

	if v.Mag() != 1 {
		t.Errorf("Error: %v", v.Mag())
	}
}

func TestMagVector(t *testing.T) {
	/* Scenario: Computing the magnitude of vector(1, 2, 3)
	   Given v ← vector(1, 2, 3)
	   Then magnitude(v) = √14 */
	v := rt.NewVector(1, 2, 3)

	if v.Mag() != math.Sqrt(14) {
		t.Errorf("Error: %v", v.Mag())
	}
}

func TestMagVectorNegative(t *testing.T) {
	/* Scenario: Computing the magnitude of vector(-1, -2, -3)
	   Given v ← vector(-1, -2, -3)
	   Then magnitude(v) = √14 */
	v := rt.NewVector(-1, -2, -3)

	if v.Mag() != math.Sqrt(14) {
		t.Errorf("Error: %v", v.Mag())
	}
}

func TestNormVectorX(t *testing.T) {
	/* Scenario: Normalizing vector(4, 0, 0) gives (1, 0, 0)
	   Given v ← vector(4, 0, 0)
	   Then normalize(v) = vector(1, 0, 0) */
	v := rt.NewVector(4, 0, 0)

	result := v.Norm()

	expected := rt.NewVector(1, 0, 0)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestNormVector(t *testing.T) {
	/* Scenario: Normalizing vector(1, 2, 3)
	   Given v ← vector(1, 2, 3)
	                                   # vector(1/√14,   2/√14,   3/√14)
	   Then normalize(v) = approximately vector(0.26726, 0.53452, 0.80178) */
	v := rt.NewVector(1, 2, 3)

	result := v.Norm()

	expected := rt.NewVector(0.26726, 0.53452, 0.80178)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
func TestMagNormVector(t *testing.T) {
	/* Scenario: The magnitude of a normalized vector
	   Given v ← vector(1, 2, 3)
	   When norm ← normalize(v)
	   Then magnitude(norm) = 1 */
	v := rt.NewVector(1, 2, 3)

	result := v.Norm().Mag()

	if result != 1 {
		t.Errorf("Error: %v", result)
	}
}

func TestDotVector(t *testing.T) {
	/* Scenario: The dot product of two tuples
	   Given a ← vector(1, 2, 3)
	     And b ← vector(2, 3, 4)
	   Then dot(a, b) = 20 */

	a := rt.NewVector(1, 2, 3)
	b := rt.NewVector(2, 3, 4)

	result := a.Dot(b)

	if result != 20 {
		t.Errorf("Error: %v", result)
	}

}

func TestCrossVector(t *testing.T) {
	/* Scenario: The cross product of two vectors
	   Given a ← vector(1, 2, 3)
	     And b ← vector(2, 3, 4)
	   Then cross(a, b) = vector(-1, 2, -1)
		 And cross(b, a) = vector(1, -2, 1) */
	a := rt.NewVector(1, 2, 3)
	b := rt.NewVector(2, 3, 4)

	result := a.Cross(b)

	expected := rt.NewVector(-1, 2, -1)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

	result = b.Cross(a)

	expected = rt.NewVector(1, -2, 1)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestReflectVector45(t *testing.T) {
	/* Scenario: Reflecting a vector approaching at 45°
	   Given v ← vector(1, -1, 0)
	     And n ← vector(0, 1, 0)
	   When r ← reflect(v, n)
	   Then r = vector(1, 1, 0) */
	v := rt.NewVector(1, -1, 0)
	n := rt.NewVector(0, 1, 0)

	r := v.Reflect(n)

	if !r.Equals(rt.NewVector(1, 1, 0)) {
		t.Errorf("Error: %v", r)
	}
}
func TestReflectVectorSlanted(t *testing.T) {
	/* Scenario: Reflecting a vector off a slanted surface
	   Given v ← vector(0, -1, 0)
	     And n ← vector(√2/2, √2/2, 0)
	   When r ← reflect(v, n)
	   Then r = vector(1, 0, 0) */
	v := rt.NewVector(0, -1, 0)
	n := rt.NewVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	r := v.Reflect(n)

	if !r.Equals(rt.NewVector(1, 0, 0)) {
		t.Errorf("Error: %v", r)
	}
}
