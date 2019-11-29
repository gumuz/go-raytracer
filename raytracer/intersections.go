package raytracer

type Intersected interface {
}

type Intersection struct {
	T      float64
	Object Intersected
}

type Intersections []*Intersection

func NewIntersections(is ...*Intersection) Intersections {
	intersections := make(Intersections, len(is))

	for idx, intersection := range is {
		intersections[idx] = intersection
	}
	return intersections
}

func NewIntersection(t float64, object Intersected) *Intersection {
	return &Intersection{t, object}
}

func (i Intersections) Hit() *Intersection {
	var hit *Intersection

	for _, intersection := range i {
		if intersection.T < 0 {
			continue
		}
		if hit == nil || intersection.T < hit.T {
			hit = intersection
			continue
		}
	}

	return hit
}
