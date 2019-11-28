package raytracer_test

import (
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
	r := rt.NewRay(rt.Point(0, 0, -5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 1, -5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 2, -5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 0, 0), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 0, 5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 0, -5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 0, -5), rt.Vector(0, 0, 1))
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
	r := rt.NewRay(rt.Point(0, 0, -5), rt.Vector(0, 0, 1))
	s := rt.NewSphere()

	s.SetTransform(rt.Translation(5, 0, 0))
	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Errorf("Error: %v", len(xs))
	}
}
