package raytracer_test

import (
	"math"
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestTranslationMatrix(t *testing.T) {
	/* Scenario: Multiplying by a translation matrix
	Given transform ← translation(5, -3, 2)
		And p ← point(-3, 4, 5)
	Then transform * p = point(2, 1, 7) */
	transform := rt.Translation(5, -3, 2)
	p := rt.Point(-3, 4, 5)

	result := transform.MulT(p)
	expected := rt.Point(2, 1, 7)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestTranslationMatrixInverse(t *testing.T) {
	/* Scenario: Multiplying by the inverse of a translation matrix
	Given transform ← translation(5, -3, 2)
		And inv ← inverse(transform)
		And p ← point(-3, 4, 5)
	Then inv * p = point(-8, 7, 3) */
	transform := rt.Translation(5, -3, 2)
	inv := transform.Inv()
	p := rt.Point(-3, 4, 5)

	result := inv.MulT(p)
	expected := rt.Point(-8, 7, 3)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestTranslationMatrixVector(t *testing.T) {
	/* Scenario: Translation does not affect vectors
	Given transform ← translation(5, -3, 2)
		And v ← vector(-3, 4, 5)
	Then transform * v = v */
	transform := rt.Translation(5, -3, 2)
	v := rt.Vector(-3, 4, 5)

	result := transform.MulT(v)

	if !result.Equals(v) {
		t.Errorf("Error: %v", result)
	}
}

func TestScalingMatrixPoint(t *testing.T) {
	/* Scenario: A scaling matrix applied to a point
	Given transform ← scaling(2, 3, 4)
	  And p ← point(-4, 6, 8)
	 Then transform * p = point(-8, 18, 32) */
	transform := rt.Scaling(2, 3, 4)
	p := rt.Point(-4, 6, 8)

	result := transform.MulT(p)
	expected := rt.Point(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestScalingMatrixVector(t *testing.T) {
	/* Scenario: A scaling matrix applied to a vector
	Given transform ← scaling(2, 3, 4)
	  And v ← vector(-4, 6, 8)
	 Then transform * v = vector(-8, 18, 32) */
	transform := rt.Scaling(2, 3, 4)
	v := rt.Vector(-4, 6, 8)

	result := transform.MulT(v)
	expected := rt.Vector(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
func TestScalingMatrixInverse(t *testing.T) {
	/* Scenario: Multiplying by the inverse of a scaling matrix
	Given transform ← scaling(2, 3, 4)
	  And inv ← inverse(transform)
	  And v ← vector(-4, 6, 8)
	 Then inv * v = vector(-2, 2, 2) */
	transform := rt.Scaling(2, 3, 4)
	inv := transform.Inv()
	v := rt.Vector(-4, 6, 8)

	result := inv.MulT(v)
	expected := rt.Vector(-2, 2, 2)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
func TestScalingMatrixReflect(t *testing.T) {
	/* Scenario: Reflection is scaling by a negative value
	Given transform ← scaling(-1, 1, 1)
	  And p ← point(2, 3, 4)
	 Then transform * p = point(-2, 3, 4) */
	transform := rt.Scaling(-1, 1, 1)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(-2, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}
func TestRotateMatrixX(t *testing.T) {
	/* Scenario: Rotating a point around the x axis
	Given p ← point(0, 1, 0)
	  And half_quarter ← rotation_x(π / 4)
	  And full_quarter ← rotation_x(π / 2)
	Then half_quarter * p = point(0, √2/2, √2/2)
	  And full_quarter * p = point(0, 0, 1) */
	p := rt.Point(0, 1, 0)

	half_quarter := rt.RotationX(math.Pi / 4)
	result := half_quarter.MulT(p)
	expected := rt.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

	full_quarter := rt.RotationX(math.Pi / 2)
	result = full_quarter.MulT(p)
	expected = rt.Point(0, 0, 1)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestRotateMatrixXReversed(t *testing.T) {
	/* Scenario: The inverse of an x-rotation rotates in the opposite direction
	Given p ← point(0, 1, 0)
	  And half_quarter ← rotation_x(π / 4)
	  And inv ← inverse(half_quarter)
	Then inv * p = point(0, √2/2, -√2/2) */
	p := rt.Point(0, 1, 0)

	half_quarter := rt.RotationX(math.Pi / 4)
	inv := half_quarter.Inv()

	result := inv.MulT(p)
	expected := rt.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

}

func TestRotateMatrixY(t *testing.T) {
	/* Scenario: Rotating a point around the y axis
	Given p ← point(0, 0, 1)
	  And half_quarter ← rotation_y(π / 4)
	  And full_quarter ← rotation_y(π / 2)
	Then half_quarter * p = point(√2/2, 0, √2/2)
	  And full_quarter * p = point(1, 0, 0) */
	p := rt.Point(0, 0, 1)

	half_quarter := rt.RotationY(math.Pi / 4)
	result := half_quarter.MulT(p)
	expected := rt.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

	full_quarter := rt.RotationY(math.Pi / 2)
	result = full_quarter.MulT(p)
	expected = rt.Point(1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestRotateMatrixZ(t *testing.T) {
	/* Scenario: Rotating a point around the z axis
	Given p ← point(0, 1, 0)
	  And half_quarter ← rotation_z(π / 4)
	  And full_quarter ← rotation_z(π / 2)
	Then half_quarter * p = point(-√2/2, √2/2, 0)
	  And full_quarter * p = point(-1, 0, 0) */
	p := rt.Point(0, 1, 0)

	half_quarter := rt.RotationZ(math.Pi / 4)
	result := half_quarter.MulT(p)
	expected := rt.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

	full_quarter := rt.RotationZ(math.Pi / 2)
	result = full_quarter.MulT(p)
	expected = rt.Point(-1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixXtoY(t *testing.T) {
	/* Scenario: A shearing transformation moves x in proportion to y
	Given transform ← shearing(1, 0, 0, 0, 0, 0)
	  And p ← point(2, 3, 4)
	Then transform * p = point(5, 3, 4) */
	transform := rt.Shearing(1, 0, 0, 0, 0, 0)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(5, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixXtoZ(t *testing.T) {
	/* Scenario: A shearing transformation moves x in proportion to z
	Given transform ← shearing(0, 1, 0, 0, 0, 0)
	  And p ← point(2, 3, 4)
	Then transform * p = point(6, 3, 4) */
	transform := rt.Shearing(0, 1, 0, 0, 0, 0)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(6, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixYtoX(t *testing.T) {
	/* Scenario: A shearing transformation moves y in proportion to x
	Given transform ← shearing(0, 0, 1, 0, 0, 0)
	  And p ← point(2, 3, 4)
	Then transform * p = point(2, 5, 4) */
	transform := rt.Shearing(0, 0, 1, 0, 0, 0)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(2, 5, 4)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixYtoZ(t *testing.T) {
	/* Scenario: A shearing transformation moves y in proportion to z
	Given transform ← shearing(0, 0, 0, 1, 0, 0)
	  And p ← point(2, 3, 4)
	Then transform * p = point(2, 7, 4) */
	transform := rt.Shearing(0, 0, 0, 1, 0, 0)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(2, 7, 4)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixZtoX(t *testing.T) {
	/* Scenario: A shearing transformation moves z in proportion to x
	Given transform ← shearing(0, 0, 0, 0, 1, 0)
	  And p ← point(2, 3, 4)
	Then transform * p = point(2, 3, 6) */
	transform := rt.Shearing(0, 0, 0, 0, 1, 0)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(2, 3, 6)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestShearMatrixZtoY(t *testing.T) {
	/* Scenario: A shearing transformation moves z in proportion to y
	Given transform ← shearing(0, 0, 0, 0, 0, 1)
	  And p ← point(2, 3, 4)
	Then transform * p = point(2, 3, 7) */
	transform := rt.Shearing(0, 0, 0, 0, 0, 1)
	p := rt.Point(2, 3, 4)

	result := transform.MulT(p)
	expected := rt.Point(2, 3, 7)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestChainTransform(t *testing.T) {
	/* Scenario: Individual transformations are applied in sequence
	Given p ← point(1, 0, 1)
	  And A ← rotation_x(π / 2)
	  And B ← scaling(5, 5, 5)
	  And C ← translation(10, 5, 7)
	# apply rotation first
	When p2 ← A * p
	Then p2 = point(1, -1, 0)
	# then apply scaling
	When p3 ← B * p2
	Then p3 = point(5, -5, 0)
	# then apply translation
	When p4 ← C * p3
	Then p4 = point(15, 0, 7) */
	p := rt.Point(1, 0, 1)
	A := rt.RotationX(math.Pi / 2)
	B := rt.Scaling(5, 5, 5)
	C := rt.Translation(10, 5, 7)

	p2 := A.MulT(p)

	if !p2.Equals(rt.Point(1, -1, 0)) {
		t.Errorf("Error: %v", p2)
	}

	p3 := B.MulT(p2)

	if !p3.Equals(rt.Point(5, -5, 0)) {
		t.Errorf("Error: %v", p3)
	}

	p4 := C.MulT(p3)

	if !p4.Equals(rt.Point(15, 0, 7)) {
		t.Errorf("Error: %v", p4)
	}
}

func TestChainTransformReversed(t *testing.T) {
	/* Scenario: Chained transformations must be applied in reverse order
	Given p ← point(1, 0, 1)
	  And A ← rotation_x(π / 2)
	  And B ← scaling(5, 5, 5)
	  And C ← translation(10, 5, 7)
	When T ← C * B * A
	Then T * p = point(15, 0, 7) */
	p := rt.Point(1, 0, 1)
	A := rt.RotationX(math.Pi / 2)
	B := rt.Scaling(5, 5, 5)
	C := rt.Translation(10, 5, 7)

	result := C.Mul(B).Mul(A).MulT(p)

	if !result.Equals(rt.Point(15, 0, 7)) {
		t.Errorf("Error: %v", result)
	}
}

func TestViewTransformDefault(t *testing.T) {
	/* Scenario: The transformation matrix for the default orientation
	Given from ← point(0, 0, 0)
	  And to ← point(0, 0, -1)
	  And up ← vector(0, 1, 0)
	When t ← view_transform(from, to, up)
	Then t = identity_matrix */
	// from := rt.Point(0, 0, 0)
	// to := rt.NewPoint(0, 0, -1)
	// up := rt.NewPoint(0, 1, 0)

	// vt := ViewTransform(from, to, up)

	// if !vt.Equals(rt.Identity) {
	// 	t.Errorf("Error: %v", vt)
	// }
}

// Scenario: A view transformation matrix looking in positive z direction
//   Given from ← point(0, 0, 0)
//     And to ← point(0, 0, 1)
//     And up ← vector(0, 1, 0)
//   When t ← view_transform(from, to, up)
//   Then t = scaling(-1, 1, -1)

// Scenario: The view transformation moves the world
//   Given from ← point(0, 0, 8)
//     And to ← point(0, 0, 0)
//     And up ← vector(0, 1, 0)
//   When t ← view_transform(from, to, up)
//   Then t = translation(0, 0, -8)

// Scenario: An arbitrary view transformation
//   Given from ← point(1, 3, 2)
//     And to ← point(4, -2, 8)
//     And up ← vector(1, 1, 0)
//   When t ← view_transform(from, to, up)
//   Then t is the following 4x4 matrix:
//       | -0.50709 | 0.50709 |  0.67612 | -2.36643 |
//       |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
//       | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
//       |  0.00000 | 0.00000 |  0.00000 |  1.00000 |
