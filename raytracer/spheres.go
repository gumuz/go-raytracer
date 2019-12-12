package raytracer

import "math"

type Sphere struct {
	Transform Matrix
	Material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{Identity(), NewMaterial()}
}

func (s *Sphere) SetTransform(transform Matrix) {
	s.Transform = transform
}

func (s *Sphere) Intersect(r *Ray) Intersections {
	transformedRay := r.Transform(s.Transform.Inv())

	sphereToRay := transformedRay.Origin.Sub(NewPoint(0, 0, 0))

	a := transformedRay.Direction.Dot(transformedRay.Direction)
	b := 2 * transformedRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	disc := math.Pow(b, 2) - 4*a*c
	if disc < 0 {
		return NewIntersections()
	}

	t1 := (-b - math.Sqrt(disc)) / (2 * a)
	t2 := (-b + math.Sqrt(disc)) / (2 * a)

	intersections := NewIntersections(
		NewIntersection(t1, s),
		NewIntersection(t2, s),
	)
	return intersections
}

func (s *Sphere) NormalAt(p *Tuple) *Tuple {
	objPoint := s.Transform.Inv().MulT(p)
	objNormal := objPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := s.Transform.Inv().Trans().MulT(objNormal)
	worldNormal.W = 0
	return worldNormal.Norm()
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}
