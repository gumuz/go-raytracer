package raytracer_test

import (
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestNewRay(t *testing.T) {
	/* Scenario: Creating and querying a ray
	Given origin ← point(1, 2, 3)
	  And direction ← vector(4, 5, 6)
	When r ← ray(origin, direction)
	Then r.origin = origin
	  And r.direction = direction */
	origin := rt.NewPoint(1, 2, 3)
	direction := rt.NewVector(4, 5, 6)
	r := &rt.Ray{origin, direction}

	if !r.Origin.Equals(origin) {
		t.Errorf("Error: %v", r.Origin)
	}

	if !r.Direction.Equals(direction) {
		t.Errorf("Error: %v", r.Origin)
	}

}

func TestRayPos(t *testing.T) {
	/* Scenario: Computing a point from a distance
	Given r ← ray(point(2, 3, 4), vector(1, 0, 0))
	Then position(r, 0) = point(2, 3, 4)
	  And position(r, 1) = point(3, 3, 4)
	  And position(r, -1) = point(1, 3, 4)
	  And position(r, 2.5) = point(4.5, 3, 4) */
	r := &rt.Ray{rt.NewPoint(2, 3, 4), rt.NewVector(1, 0, 0)}

	if !r.Pos(0).Equals(rt.NewPoint(2, 3, 4)) {
		t.Errorf("Error: %v", r.Pos(0))
	}

	if !r.Pos(1).Equals(rt.NewPoint(3, 3, 4)) {
		t.Errorf("%v is incorrect", r.Pos(1))
	}

	if !r.Pos(-1).Equals(rt.NewPoint(1, 3, 4)) {
		t.Errorf("%v is incorrect", r.Pos(-1))
	}

	if !r.Pos(2.5).Equals(rt.NewPoint(4.5, 3, 4)) {
		t.Errorf("%v is incorrect", r.Pos(2.5))
	}
}

func TestRayTranslation(t *testing.T) {
	/* Scenario: Translating a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	  And m ← translation(3, 4, 5)
	When r2 ← transform(r, m)
	Then r2.origin = point(4, 6, 8)
	  And r2.direction = vector(0, 1, 0) */
	r := &rt.Ray{rt.NewPoint(1, 2, 3), rt.NewVector(0, 1, 0)}
	m := rt.Translation(3, 4, 5)

	r2 := r.Transform(m)

	if !r2.Origin.Equals(rt.NewPoint(4, 6, 8)) {
		t.Errorf("Error: %v", r2.Origin)
	}

	if !r2.Direction.Equals(rt.NewVector(0, 1, 0)) {
		t.Errorf("Error: %v", r2.Origin)
	}
}

func TestRayScaling(t *testing.T) {
	/* Scenario: Scaling a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	  And m ← scaling(2, 3, 4)
	When r2 ← transform(r, m)
	Then r2.origin = point(2, 6, 12)
	  And r2.direction = vector(0, 3, 0) */
	r := &rt.Ray{rt.NewPoint(1, 2, 3), rt.NewVector(0, 1, 0)}
	m := rt.Scaling(2, 3, 4)

	r2 := r.Transform(m)

	if !r2.Origin.Equals(rt.NewPoint(2, 6, 12)) {
		t.Errorf("Error: %v", r2.Origin)
	}

	if !r2.Direction.Equals(rt.NewVector(0, 3, 0)) {
		t.Errorf("Error: %v", r2.Origin)
	}
}
