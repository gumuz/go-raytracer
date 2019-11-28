package raytracer_test

import (
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestNewIntersection(t *testing.T) {
	/* Scenario: An intersection encapsulates t and object
	   Given s ← sphere()
	   When i ← intersection(3.5, s)
	   Then i.t = 3.5
	     And i.object = s */
	s := rt.NewSphere()

	i := rt.NewIntersection(3.5, s)

	if i.T != 3.5 {
		t.Errorf("Error: %v", i.T)
	}

	if i.Object != s {
		t.Errorf("Error: %v", i.Object)
	}
}
func TestNewIntersections(t *testing.T) {
	/* Scenario: Aggregating intersections
	   Given s ← sphere()
	     And i1 ← intersection(1, s)
	     And i2 ← intersection(2, s)
	   When xs ← intersections(i1, i2)
	   Then xs.count = 2
	     And xs[0].t = 1
		 And xs[1].t = 2 */
	s := rt.NewSphere()
	i1 := rt.NewIntersection(1, s)
	i2 := rt.NewIntersection(2, s)

	xs := rt.Intersections{i1, i2}

	if len(xs) != 2 {
		t.Errorf("Error: %v", len(xs))
	}

	if xs[0].T != 1 {
		t.Errorf("Error: %v", xs[0].T)
	}

	if xs[1].T != 2 {
		t.Errorf("Error: %v", xs[1].T)
	}
}

func TestIntersectionsAllPositive(t *testing.T) {
	/* Scenario: The hit, when all intersections have positive t
	   Given s ← sphere()
	     And i1 ← intersection(1, s)
	     And i2 ← intersection(2, s)
	     And xs ← intersections(i2, i1)
	   When i ← hit(xs)
	   Then i = i1 */
	s := rt.NewSphere()
	i1 := rt.NewIntersection(1, s)
	i2 := rt.NewIntersection(2, s)
	xs := rt.Intersections{i1, i2}

	i := xs.Hit()

	if i != i1 {
		t.Errorf("Error: %v", i)
	}
}
func TestIntersectionsSomePositive(t *testing.T) {
	/* Scenario: The hit, when some intersections have negative t
	   Given s ← sphere()
	     And i1 ← intersection(-1, s)
	     And i2 ← intersection(1, s)
	     And xs ← intersections(i2, i1)
	   When i ← hit(xs)
	   Then i = i2 */
	s := rt.NewSphere()
	i1 := rt.NewIntersection(-1, s)
	i2 := rt.NewIntersection(1, s)
	xs := rt.Intersections{i1, i2}

	i := xs.Hit()

	if i != i2 {
		t.Errorf("Error: %v", i)
	}
}
func TestIntersectionsAllNegative(t *testing.T) {
	/* Scenario: The hit, when all intersections have negative t
	   Given s ← sphere()
	     And i1 ← intersection(-2, s)
	     And i2 ← intersection(-1, s)
	     And xs ← intersections(i2, i1)
	   When i ← hit(xs)
	   Then i is nothing */
	s := rt.NewSphere()
	i1 := rt.NewIntersection(-2, s)
	i2 := rt.NewIntersection(-1, s)
	xs := rt.Intersections{i1, i2}

	i := xs.Hit()

	if i != nil {
		t.Errorf("Error: %v", i)
	}
}

func TestIntersectionsLowestPositive(t *testing.T) {
	/* Scenario: The hit is always the lowest nonnegative intersection
	     Given s ← sphere()
	     And i1 ← intersection(5, s)
	     And i2 ← intersection(7, s)
	     And i3 ← intersection(-3, s)
	     And i4 ← intersection(2, s)
	     And xs ← intersections(i1, i2, i3, i4)
	   When i ← hit(xs)
	   Then i = i4 */
	s := rt.NewSphere()
	i1 := rt.NewIntersection(5, s)
	i2 := rt.NewIntersection(7, s)
	i3 := rt.NewIntersection(-3, s)
	i4 := rt.NewIntersection(2, s)
	xs := rt.Intersections{i1, i2, i3, i4}

	i := xs.Hit()

	if i != i4 {
		t.Errorf("Error: %v", i)
	}
}
