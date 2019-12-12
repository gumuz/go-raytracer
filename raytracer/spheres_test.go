package raytracer_test

import (
	"math"
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestRayIntersect(t *testing.T) {
	/* Scenario: A ray intersects a sphere at two points
	   Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0] = 4.0
	 	And xs[1] = 6.0 */
	r := rt.NewRay(rt.NewPoint(0, 0, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != 4 {
		t.Errorf("Error: %v", xs[0])
	}

	if xs[1].T != 6 {
		t.Errorf("Error: %v", xs[1])
	}
}

func TestRayIntersectTangent(t *testing.T) {
	/* Scenario: A ray intersects a sphere at a tangent
	   Given r ← ray(point(0, 1, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0] = 5.0
	 	And xs[1] = 5.0 */
	r := rt.NewRay(rt.NewPoint(0, 1, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != 5 {
		t.Errorf("Error: %v", xs[0].T)
	}

	if xs[1].T != 5 {
		t.Errorf("Error: %v", xs[1].T)
	}
}

func TestRayNoIntersect(t *testing.T) {
	/* Scenario: A ray misses a sphere
	   Given r ← ray(point(0, 2, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 0 */
	r := rt.NewRay(rt.NewPoint(0, 2, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Errorf("Error: %v", len(xs))
	}
}

func TestRayIntersectInside(t *testing.T) {
	/* Scenario: A ray originates inside a sphere
	   Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0] = -1.0
	 	And xs[1] = 1.0 */
	r := rt.NewRay(rt.NewPoint(0, 0, 0), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != -1 {
		t.Errorf("Error: %v", xs[0].T)
	}

	if xs[1].T != 1 {
		t.Errorf("Error: %v", xs[1].T)
	}
}

func TestRayIntersectBehind(t *testing.T) {
	/* Scenario: A sphere is behind a ray
	   Given r ← ray(point(0, 0, 5), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0] = -6.0
	 	And xs[1] = -4.0 */
	r := rt.NewRay(rt.NewPoint(0, 0, 5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != -6 {
		t.Errorf("Error: %v", xs[0].T)
	}

	if xs[1].T != -4 {
		t.Errorf("Error: %v", xs[1].T)
	}
}

func TestRayIntersectSetsObject(t *testing.T) {
	/* Scenario: Intersect sets the object on the intersection
	   Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0].object = s
		 And xs[1].object = s */
	r := rt.NewRay(rt.NewPoint(0, 0, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].Object != s {
		t.Errorf("Error: %v", xs[0].Object)
	}

	if xs[1].Object != s {
		t.Errorf("Error: %v", xs[1].Object)
	}
}

func TestSphereDefaultTransformation(t *testing.T) {
	/* Scenario: A sphere's default transformation
	   Given s ← sphere()
	   Then s.transform = identity_matrix */
	s := rt.NewSphere()

	if !s.Transform.Equals(rt.Identity()) {
		t.Errorf("Error: %v", s.Transform)
	}
}

func TestSphereChangeTransformation(t *testing.T) {
	/* Scenario: Changing a sphere's transformation
	   Given s ← sphere()
	     And t ← translation(2, 3, 4)
	   When set_transform(s, t)
	   Then s.transform = t */
	s := rt.NewSphere()
	tr := rt.Translation(2, 3, 4)

	s.SetTransform(tr)

	if !s.Transform.Equals(tr) {
		t.Errorf("Error: %v", s.Transform)
	}
}

func TestRayIntersectScaledSphere(t *testing.T) {
	/* Scenario: Intersecting a scaled sphere with a ray
	   Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When set_transform(s, scaling(2, 2, 2))
	     And xs ← intersect(s, r)
	   Then xs.count = 2
	     And xs[0].t = 3
		 And xs[1].t = 7 */
	r := rt.NewRay(rt.NewPoint(0, 0, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	s.SetTransform(rt.Scaling(2, 2, 2))
	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != 3 {
		t.Errorf("Error: %v", xs[0].T)
	}

	if xs[1].T != 7 {
		t.Errorf("Error: %v", xs[1].T)
	}
}

func TestRayIntersectTranslatedSphere(t *testing.T) {
	/* Scenario: Intersecting a translated sphere with a ray
	   Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	     And s ← sphere()
	   When set_transform(s, translation(5, 0, 0))
	     And xs ← intersect(s, r)
	   Then xs.count = 0 */
	r := rt.NewRay(rt.NewPoint(0, 0, -5), rt.NewVector(0, 0, 1))
	s := rt.NewSphere()

	s.SetTransform(rt.Translation(5, 0, 0))
	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Errorf("Error: %v", len(xs))
	}
}

func TestNormalX(t *testing.T) {
	/* Scenario: The normal on a sphere at a point on the x axis
	   Given s ← sphere()
	   When n ← normal_at(s, point(1, 0, 0))
	   Then n = vector(1, 0, 0) */
	s := rt.NewSphere()

	n := s.NormalAt(rt.NewPoint(1, 0, 0))

	if !n.Equals(rt.NewVector(1, 0, 0)) {
		t.Errorf("Error: %v", n)
	}
}
func TestNormalY(t *testing.T) {
	/* Scenario: The normal on a sphere at a point on the y axis
	   Given s ← sphere()
	   When n ← normal_at(s, point(0, 1, 0))
	   Then n = vector(0, 1, 0) */
	s := rt.NewSphere()

	n := s.NormalAt(rt.NewPoint(0, 1, 0))

	if !n.Equals(rt.NewVector(0, 1, 0)) {
		t.Errorf("Error: %v", n)
	}
}
func TestNormalZ(t *testing.T) {
	/* Scenario: The normal on a sphere at a point on the z axis
	   Given s ← sphere()
	   When n ← normal_at(s, point(0, 0, 1))
	   Then n = vector(0, 0, 1) */
	s := rt.NewSphere()

	n := s.NormalAt(rt.NewPoint(0, 0, 1))

	if !n.Equals(rt.NewVector(0, 0, 1)) {
		t.Errorf("Error: %v", n)
	}
}
func TestNormalNonAxial(t *testing.T) {
	/* Scenario: The normal on a sphere at a nonaxial point
	   Given s ← sphere()
	   When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	   Then n = vector(√3/3, √3/3, √3/3) */
	s := rt.NewSphere()

	n := s.NormalAt(rt.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(rt.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)) {
		t.Errorf("Error: %v", n)
	}
}
func TestNormalIsNormalizedVector(t *testing.T) {
	/* Scenario: The normal is a normalized vector
	   Given s ← sphere()
	   When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	   Then n = normalize(n) */
	s := rt.NewSphere()

	n := s.NormalAt(rt.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(n.Norm()) {
		t.Errorf("Error: %v", n)
	}
}

func TestNormalOnTranslatedSphere(t *testing.T) {
	/* Scenario: Computing the normal on a translated sphere
	   Given s ← sphere()
	     And set_transform(s, translation(0, 1, 0))
	   When n ← normal_at(s, point(0, 1.70711, -0.70711))
	   Then n = vector(0, 0.70711, -0.70711) */
	s := rt.NewSphere()
	s.SetTransform(rt.Translation(0, 1, 0))

	n := s.NormalAt(rt.NewPoint(0, 1.70711, -0.70711))

	if !n.Equals(rt.NewVector(0, 0.70711, -0.70711)) {
		t.Errorf("Error: %v", n)
	}
}

func TestNormalOnTransformedSphere(t *testing.T) {
	/* Scenario: Computing the normal on a transformed sphere
	   Given s ← sphere()
	     And m ← scaling(1, 0.5, 1) * rotation_z(π/5)
	     And set_transform(s, m)
	   When n ← normal_at(s, point(0, √2/2, -√2/2))
	   Then n = vector(0, 0.97014, -0.24254) */
	s := rt.NewSphere()
	s.SetTransform(rt.Scaling(1, 0.5, 1).Mul(rt.RotationZ(math.Pi / 5)))

	n := s.NormalAt(rt.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	if !n.Equals(rt.NewVector(0, 0.97014, -0.24254)) {
		t.Errorf("Error: %v", n)
	}
}

func TestSphereDefaultMaterial(t *testing.T) {
	/* Scenario: A sphere has a default material
	   Given s ← sphere()
	   When m ← s.material
	   Then m = material() */
	s := rt.NewSphere()

	if !s.Material.Equals(rt.NewMaterial()) {
		t.Errorf("Error: %v", s.Material)
	}
}

func TestSphereAssignedMaterial(t *testing.T) {
	/* Scenario: A sphere may be assigned a material
	   Given s ← sphere()
	     And m ← material()
	     And m.ambient ← 1
	   When s.material ← m
	   Then s.material = m */
	s := rt.NewSphere()
	m := rt.NewMaterial()
	m.Ambient = 1
	s.Material = m

	if !s.Material.Equals(m) {
		t.Errorf("Error: %v", s.Material)
	}
}
